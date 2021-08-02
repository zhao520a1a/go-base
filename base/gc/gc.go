package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	//获取已分配内存的总量，单位是 Kb
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)

	//显式的触发 GC
	//runtime.GC()

	//如果一个对象 obj需要在被从内存移除前执行一些特殊操作，比如写到日志文件中
	//runtime.SetFinalizer(obj, func(obj *typeObj))
}
