package cond

import (
	"sync"
	"testing"
)

func TestCondSignal(t *testing.T) {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 2
	running := make(chan bool, n)
	awake := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			running <- true
			c.Wait()
			awake <- true
			m.Unlock()
		}()
	}
	for i := 0; i < n; i++ {
		<-running // Wait for everyone to run.
	}
	for n > 0 {
		select {
		case <-awake:
			t.Fatal("goroutine not asleep")
		default:
		}
		m.Lock()
		c.Signal()
		m.Unlock()
		<-awake // Will deadlock if no goroutine wakes up
		select {
		case <-awake:
			t.Fatal("too many goroutines awake")
		default:
		}
		n--
	}
	c.Signal()
}
