package example

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
分别用 chan、WaitGroup、Context 来协调协程间调用
*/

func TestCoordinate(t *testing.T) {
	coordinateWithChan()
	fmt.Println()
	coordinateWithWaitGroup()
	fmt.Println()
	coordinateWithContext()
}

func coordinateWithChan() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	sign := make(chan struct{}, 2)
	for i := 1; i <= total; i++ {
		go addNum(&num, 1, func() {
			sign <- struct{}{}
		})
	}

	for i := 1; i <= total; i++ {
		<-sign
	}
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("End.")
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	<-cxt.Done()
}

// addNum 用于原子地增加一次numP所指的变量的值。
func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		}
	}
}
