/*
下面的程序验证, 由于缓存不一致导致性能下降的情况以及解决方案。

我们一个接一个地实例化两个结构体。因此，这些结构体应在内存中连续分配。然后，我们创建两个 goroutine，其中每个 goroutine 都访问其结构：
分析：第二个结构的变量 n 仅由第二个 goroutine 访问。但是，由于结构在内存中是连续的，因此它会出现在两条高速缓存行中

(前提假设两个 goroutine 分别被单独内核上的线程调度；实际上不一定是这种情况，因为Go 1.5 版本之前，默认使用的是单核心执行。从 Go 1.5 版本开始，默认执行 runtime.GOMAXPROCS(runtime.NumCPU()) 语句以便让代码并发执行，最大效率地利用 CPU。）
*/
package main

import (
	"sync"
	"testing"
)

type SimpleStruct struct {
	n int
}

const M = 10000000 //一千万

func BenchmarkStructureFalseSharing(b *testing.B) {
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

type PaddedStruct struct {
	_ CacheLinePad //内存空间填充
	n int
	_ CacheLinePad //内存空间填充
}

type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

const CacheLinePadSize = 64

func BenchmarkStructurePadding(b *testing.B) {
	structA := PaddedStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
