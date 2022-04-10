package currency

import "sync"

/*
通过使用 sync 包可以解决同一时间只能一个线程访问变量或 map-demo 类型数据的问题。当然也通过 goroutines 和 channels 来解决问题
*/

type Info struct {
	mu   sync.Mutex   //互斥锁
	rmu  sync.RWMutex // 读写锁
	once sync.Once
	//....
}

func Update1(info *Info) {
	info.mu.Lock() //确保同一时间只能有一个线程进入临界区

	//... 更新变量

	info.mu.Unlock()
}

func Update2(info *Info) {
	//info.rmu.Lock() //和普通的 Mutex 作用相同
	//info.rmu.Unlock() //和普通的 Mutex 作用相同
	info.rmu.RLock() //允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作

	//... 更新变量

	info.rmu.RUnlock()
}

func limitOnceCall(info *Info, f func()) {
	info.once.Do(f) //确保被调用函数只能被调用一次。
}
