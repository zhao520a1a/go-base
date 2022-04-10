package cond

import (
	"fmt"
	"sync"
	"time"
)

// 分别通过 Channel 和 Cond 实现多协程间单向通知功能

func FunByChannel() {
	done := make(chan int, 1)

	go func() {
		time.Sleep(5 * time.Second)
		done <- 1
	}()

	fmt.Println("waiting")
	<-done
	fmt.Println("done")
}

func FunByCond() {
	cond := sync.NewCond(&sync.Mutex{})
	var flag bool
	go func() {
		time.Sleep(time.Second * 5)
		cond.L.Lock()
		flag = true
		cond.Signal()
		cond.L.Unlock()
	}()

	fmt.Println("waiting")
	cond.L.Lock()
	for !flag {
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("done")
}
