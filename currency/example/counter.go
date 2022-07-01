package example

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	cmap "github.com/orcaman/concurrent-map"
)

type CounterMapManager struct {
	mu      sync.Mutex
	dataMap cmap.ConcurrentMap
}

func NewCounterMapManager() *CounterMapManager {
	return &CounterMapManager{
		dataMap: cmap.New(),
	}
}

func (c *CounterMapManager) Set(key string, value *int64) {
	c.dataMap.Set(key, value)
}

func (c *CounterMapManager) IncrBy(key string, delta int64) {
	// fast-path
	value, ok := c.Get(key)
	if ok {
		atomic.AddInt64(value, delta)
		return
	}
	c.NewKey(key, delta)
	return
}

func (c *CounterMapManager) NewKey(key string, delta int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// slow-path
	value, ok := c.Get(key)
	if ok {
		atomic.AddInt64(value, delta)
		return
	}
	c.Set(key, &delta)
	return
}

func (c *CounterMapManager) Get(key string) (*int64, bool) {
	var item *int64
	v, ok := c.dataMap.Get(key)
	if ok {
		item, ok = v.(*int64)
		return item, ok
	}
	return item, ok
}

func (c *CounterMapManager) Reload(ctx context.Context) {
	var successNum, failNum int64
	successDataMap := make(map[string]int64)
	c.dataMap.IterCb(func(key string, v interface{}) {
		count, ok := v.(int64)
		if ok {
			fmt.Printf("key %s val %d", key, count)
			// 持久化归档
		}
	})
	for key, value := range successDataMap {
		v, ok := c.dataMap.Get(key)
		if !ok {
			continue
		}
		newVal, ok := v.(int64)
		if !ok {
			continue
		}
		if value != newVal {
			fmt.Printf("key %s oldVal %d newVal %d not equal", key, value, newVal)
			continue
		}
		c.dataMap.Remove(key)
	}
	fmt.Printf("CounterMapManager Reload finished, success %d fail %d", successNum, failNum)
	return
}
