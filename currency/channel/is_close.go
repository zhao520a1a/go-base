/*
协程的同步： 关闭通道-测试阻塞通道
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	go sendData(ch)
	getData(ch)
}

func getData(ch chan int) {

	//使用for range读channel
	//for i := range ch {
	//	fmt.Println(i)
	//}

	//使用_,ok判断channel是否关闭
	for {
		if i, ok := <-ch; ok {
			fmt.Println(i)
		} else {
			fmt.Println("receive close chan msg")
			break
		}
	}

	// 不判断 chanel 关闭会死循环
	//for {
	//	fmt.Println(<-ch)
	//}

}

func sendData(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println(isClosed(ch))
	close(ch)
	fmt.Println(isClosed(ch))

}

func isClosed(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
