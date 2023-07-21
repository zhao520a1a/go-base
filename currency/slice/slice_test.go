package slice

import (
	"fmt"
	"sync"
	"testing"
)

/*
并发访问 slice 是不安全的！
真实的输出并没有达到我们的预期，len(slice) < n。 问题出在哪？我们都知道slice是对数组一个连续片段的引用，当slice长度增加的时候，可能底层的数组会被换掉。当出在换底层数组之前，切片同时被多个goroutine拿到，并执行append操作。那么很多goroutine的append结果会被覆盖，导致n个gouroutine append后，长度小于n。
*/
func TestSyncSlice(t *testing.T) {
	var (
		slc = []int{}
		n   = 10000
		wg  sync.WaitGroup
	)

	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			slc = append(slc, i)
		}()
	}
	wg.Wait()

	fmt.Println("len:", len(slc))
	fmt.Println("done")
}

// 使用 channel 串行化操作保证并发安全
func Benchmark_Slick_Channel(b *testing.B) {
	var (
		wg sync.WaitGroup
		n  = 1000000
	)
	c := make(chan struct{})

	// new 了这个 job 后，该 job 就开始准备从 channel 接收数据了
	s := NewScheduleJob(n, func() { c <- struct{}{} })

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			s.AddData(v)
		}(i)
	}

	wg.Wait()
	s.Close()
	<-c

	fmt.Println(len(s.data))
}

// 优点是比较简单，性能相对差些
func Benchmark_Slick_Lock(b *testing.B) {
	n := 1000000
	slc := make([]int, 0, n)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()
			// 加🔐
			lock.Lock()
			defer lock.Unlock()
			slc = append(slc, a)
		}(i)

	}
	wg.Wait()
	fmt.Println(len(slc))
}

// Benchmark_Slick_Channel-12    	1000000000	         0.215 ns/op
// PASS
// Benchmark_Slick_Lock-12    	1000000000	         0.310 ns/op
// PASS
