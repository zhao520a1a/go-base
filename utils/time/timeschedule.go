package time

import (
	"context"
	"fmt"
	"log"
	"time"
)

// StartByTimePeriod 定时：时间段处理一次
func StartByTimePeriod(ctx context.Context, after time.Duration, targetFun func(context.Context) error) {
	fun := "StartByTimePeriod -->"

TimedLoop:
	for {
		err := targetFun(ctx)
		if err != nil {
			log.Printf("%s call target fun err:%v", fun, err)
		}

		select {
		case <-time.After(after):
			log.Printf("%s start next round", fun)
		case <-ctx.Done():
			log.Printf("%s about to exit", fun)
			break TimedLoop
		}
	}
}

// StartByTimePeriod1 使用 Tick/Sleep 每隔100毫秒打印“Hello TigerwolfC”
func StartByTimePeriod1(ctx context.Context) {
	// 方式1
	for range time.Tick(time.Millisecond * 100) {
		fmt.Println("Hello TigerwolfC")
	}
	// 方式2
	ticker := time.NewTicker(time.Millisecond * 100)
	for range ticker.C {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
	// 方式3
	for {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Hello TigerwolfC")
	}
}

// StartByTimeDot 定时：每天固定时间点处理一次
func StartByTimeDot(ctx context.Context, hourDot int, targetFun func(context.Context) error) {
	fun := "Manager.StartSync -->"

TimedLoop:
	for {
		err := targetFun(ctx)
		if err != nil {
			log.Printf("call target fun err:%v", err)
		}

		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), hourDot, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))

		select {
		case <-t.C:
			log.Printf("%s start next round", fun)
		case <-ctx.Done():
			log.Printf("%s about to exit", fun)
			break TimedLoop
		}
	}
}
