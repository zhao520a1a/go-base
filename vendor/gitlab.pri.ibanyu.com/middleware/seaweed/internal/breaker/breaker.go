package breaker

import (
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xutil"
)

const (
	checkTick          = time.Millisecond * 25
	defaultGranularity = time.Second * 1
	defaultThreshold   = 10
	defaultBreakerGap  = 10 // 单位: seconds
)

// TOOD 简单计数法实现熔断操作，后续改为滑动窗口或三方组件的方式
type BreakerManager struct {
	lock     sync.Mutex
	Breakers map[string]*Breaker
}

type Breaker struct {
	Rejected      int32
	RejectedStart int64
	Count         int32
}

var bm *BreakerManager

// StatBreaker state errors for breaker
func StatBreaker(cluster, table string, err error) {
	if err != nil && (strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "invalid connection")) {
		key := xutil.Concat(cluster, "_", table)
		bm.lock.Lock()
		if _, ok := bm.Breakers[key]; !ok {
			breaker := new(Breaker)
			breaker.run()
			bm.Breakers[key] = breaker
		}
		breaker := bm.Breakers[key]
		bm.lock.Unlock()
		atomic.AddInt32(&breaker.Count, 1)
	}
}

// Entry check if allow request
func Entry(cluster, table string) bool {
	key := xutil.Concat(cluster, "_", table)
	bm.lock.Lock()
	breaker := bm.Breakers[key]
	bm.lock.Unlock()
	if breaker != nil {
		return atomic.LoadInt32(&breaker.Rejected) != 1
	}
	return true
}

func (breaker *Breaker) run() {
	go func() {
		granularityTickC := time.Tick(defaultGranularity)
		checkTickC := time.Tick(checkTick)
		for {
			select {
			case <-granularityTickC:
				atomic.StoreInt32(&breaker.Count, 0)
				// check 1s/checkTick times in 1s
			case <-checkTickC:
				threshold := defaultThreshold
				breakerGap := defaultBreakerGap
				if atomic.LoadInt32(&breaker.Count) > int32(threshold) {
					atomic.StoreInt32(&breaker.Rejected, 1)
					breaker.RejectedStart = time.Now().Unix()
				} else {
					now := time.Now().Unix()
					if now-breaker.RejectedStart > int64(breakerGap) {
						atomic.StoreInt32(&breaker.Rejected, 0)
					}
				}
			}
		}
	}()
}

func init() {
	bm = &BreakerManager{Breakers: make(map[string]*Breaker)}
}
