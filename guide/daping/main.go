package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Obj struct {
	Value string
}

func (o *Obj) Incr(input *int64) {
	atomic.AddInt64(input, 1)
}

func NewObj(value string) *Obj {
	return &Obj{Value: value}
}

var global = NewObj("")

var total = 1000000

const tt = 0 * time.Millisecond

func A() {
	var a int64
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			global.Incr(&a)
			time.Sleep(tt)
			wg.Done()
		}()
	}
	wg.Wait()
}

func B() {
	var a int64
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			obj := NewObj("B")
			obj.Incr(&a)
			time.Sleep(tt)
			wg.Done()
		}()
	}
	wg.Wait()
}

func B1() {
	var a int64
	obj := NewObj("B")
	wg := sync.WaitGroup{}
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			obj.Incr(&a)
			time.Sleep(tt)
			wg.Done()
		}()
	}
	wg.Wait()
}

var p = sync.Pool{
	New: func() interface{} {
		return NewObj("aaa")
	},
}

func C() {
	var a int64
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			obj := p.Get().(*Obj)
			obj.Incr(&a)
			p.Put(obj)
			time.Sleep(tt)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	fmt.Println(time.Unix(0, 0))

	fmt.Println(testing.AllocsPerRun(1000, A))
	fmt.Println(testing.AllocsPerRun(1000, B))

}
