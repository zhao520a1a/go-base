package pattern

import "sync"

//Go 单例模式实现
var ins *singleton
var once sync.Once

type singleton struct {
}

func GetSingleton() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
