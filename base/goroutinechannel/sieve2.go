/*
打印了输出的素数，使用选择器作为它的算法。每个 prime 都有一个选择器，

函数 sieve、generate 和 filter 都是工厂；它们创建通道并返回，而且使用了协程的 lambda 函数。main 函数现在短小清晰：它调用 sieve() 返回了包含素数的通道，然后通过 fmt.Println(<-primes) 打印出来。
*/
package main

import (
	"fmt"
)

func main() {
	primes := sieve2()
	for {
		fmt.Println(<-primes)
	}

}

func sieve2() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			out <- prime
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

//将2，3，4... 发送并返回一个channel
func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//过滤掉可以被prime整除的数字，将剩余的发送并返回一个channel
func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
