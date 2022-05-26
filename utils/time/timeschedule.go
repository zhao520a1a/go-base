package time

import (
	"context"
	"log"
	"time"
)

// StartByTimePeriod 定时：时间段处理一次
func StartByTimePeriod(ctx context.Context, after time.Duration, targetFun func(context.Context) error) {
	fun := "StartByTimePeriod -->"

	// 方式1
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

	// 方式2
	for range time.Tick(after) {
		err := targetFun(ctx)
		if err != nil {
			log.Printf("%s call target fun err:%v", fun, err)
		}
	}
	// 方式3
	ticker := time.NewTicker(after)
	for range ticker.C {
		err := targetFun(ctx)
		if err != nil {
			log.Printf("%s call target fun err:%v", fun, err)
		}
	}
	// 方式4
	for {
		time.Sleep(after)
		err := targetFun(ctx)
		if err != nil {
			log.Printf("%s call target fun err:%v", fun, err)
		}
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
