package design_model

import (
	"sync"
)

//Once类型。它能保证某个操作仅且只执行一次。使用sync.Once包是安全地实现此目标的首选方式.
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
