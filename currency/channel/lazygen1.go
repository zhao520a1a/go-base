/*
简单的惰性生成器的实现
注意：从通道读取的值可能会是稍早前产生的，并不是在程序被调用时生成的。
*/
package main

import "fmt"

var resume chan int

func main() {
	resume = integers()

	// 读取数据
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}

// 提前生成数据，等待被读取
func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}
