/*
打印了输出的prime（素数），使用选择器作为它的算法。每个 prime 都有一个选择器，
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	sieve1()
	time.Sleep(3 * time.Second)
}

func sieve1() {
	ch := make(chan int)
	go generate1(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter1(ch, ch1, prime)
		ch = ch1
	}
}

//将2，3，4... 发送并返回一个channel
func generate1(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 拷贝整数到输出通道，过滤掉可以被 prime 整除的数字
func filter1(in, out chan int, prime int) {
	for {
		if i := <-in; i%prime != 0 {
			out <- i
		}
	}
}
