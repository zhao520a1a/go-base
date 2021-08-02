package time

import (
	"log"
	"time"
)

/* timeoutWarning
tag、detailed 表示超时发生位置的两个字符串参数。
start 程序开始执行的时间
timeLimit 函数执行超时阀值，单位是秒。
*/
func timeoutWarning(tag, detailed string, start time.Time, timeLimit float64) {
	tc := time.Since(start).Seconds()
	if tc > timeLimit {
		log.Println(tag, " detailed:", detailed, "timeLimit:", timeLimit, "timeoutWarning using", tc, "s")
	}
	log.Println(tag, " detailed:", detailed, "TimeRecord using", tc, "s")
}

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		log.Printf("time cost = %v\n", tc)
	}
}

// CostTimeDecorator 这个函数的执行顺序很值得去研究
func CostTimeDecorator(f func()) func() {
	return func() {
		defer timeCost()()
		defer timeoutWarning("SaveAppLogMain", "Total", time.Now(), float64(2))
		f()
	}
}

/*
Recorder
目标：
1.实现细节要剥离：时间统计实现的细节不期望在显式的写在主逻辑中。因为主逻辑中的其他逻辑和时间统计的抽象层次不在同一个层级
2.用于时间统计的代码可复用
3.统计出来的时间结果是可被处理的。
4. 对并发编程友好
*/
type Recorder interface {
	SetCost(time.Duration)
	GetCost() time.Duration
}

func RecordExeTime(rec Recorder, f func()) func() {
	return func() {
		start := time.Now()
		f()
		timeCost := time.Since(start)
		rec.SetCost(timeCost)
	}
}

func NewTimeRecorder() Recorder {
	return &timeRecorder{}
}

type timeRecorder struct {
	cost time.Duration
}

func (tr *timeRecorder) SetCost(cost time.Duration) {
	tr.cost = cost
}

func (tr *timeRecorder) GetCost() time.Duration {
	return tr.cost
}
