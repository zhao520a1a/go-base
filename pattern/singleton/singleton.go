package singleton

import "sync"

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
