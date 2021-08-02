/*
实现多路复用，多个请求共用几个协程;
*/
package main

import (
	"fmt"
)

type Request struct {
	a, b   int
	replyc chan int //请求中的回复通道
}

type binOp func(a, b int) int

func run1(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

//方式一：存在的问题： 该程序就会无限地创建Go协程，所有协程同时运行，不断消耗资源。
func server1(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run1(op, req)
		case <-quit:
			return
		}
	}
}

const MaxOutstanding = 100

var sem = make(chan int, MaxOutstanding)

//方式2：存在的问题：用semaphore来限制资源访问，只有 MaxOutstanding 个 Go 协程能同时运行；但还是为每个进入的请求都创建了新的 Go 协程。若请求来得很快，该程序就会无限地消耗资源。
func server2(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run2(op, req)
		case <-quit:
			return
		}
	}
}
func run2(op binOp, req *Request) {
	sem <- 1
	req.replyc <- op(req.a, req.b)
	<-sem
}

//方式3：为了解决每个进入的请求都创建了新的Go协程,将sem <- 1提出来，从而限制创建Go协程;
func server3(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			sem <- 1
			go run3(op, req)
		case <-quit:
			return
		}
	}
}
func run3(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
	<-sem
}

func startServer(op binOp) (reqChan chan *Request, quit chan bool) {
	reqChan = make(chan *Request)
	quit = make(chan bool)
	go server3(op, reqChan, quit)
	return reqChan, quit
}

func stopServe(quick chan bool) {
	quick <- true //通过发信号通知关闭服务器
}

func main() {
	adder, quit := startServer(func(a, b int) int {
		return a + b
	})

	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	//校验
	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at ", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

	fmt.Print("done")

	stopServe(quit)
}
