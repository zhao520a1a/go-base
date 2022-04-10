/*
用信号量是实现互斥锁，限制对资源的访问，解决读写问题；
实现类似：实现信号量的 sync 的 Go 包功能

通道的容量= 要同步的资源容量
通道的长度= 与当前资源被使用的数量
未处理的资源个数 = 容量-长度
*/

package main

type Empty interface {
}

type semaphore chan Empty

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes 互斥 */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

func main() {
	//初始化信号量
	sem := make(semaphore, 1)
	sem.Lock()
	//doSomething
	sem.Unlock()

}
