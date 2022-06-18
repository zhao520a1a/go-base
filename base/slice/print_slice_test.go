package slicedemo

import (
	"fmt"
	"testing"
)

/*
执行结果：
[10 20 30]
[]
[10 20 30 0 0 0 0 0 0 0]

问题1：为什么打印 sl[:] 时，结果为空。但打印 sl[:10] 时，结果包含了 10 个元素，还包含了函数闭包中插入的 10, 20, 30，之间有什么关系？

关键点：在 Go 语言中，本质上只有值传递！！
如果传过去的值是指向内存空间的地址，是可以对这块内存空间做修改的，实质上在调用 appendFunc(sl) 函数时，实际上修改了底层所指向的数组，即：sl 和 s 底层都是同一个数组
*/

var appendFunc = func(s []int) {
	s = append(s, 10, 20, 30)
	fmt.Println(s)
}

func TestPrintSlice(t *testing.T) {
	sl := make([]int, 0, 10) //  len=0 cap=10
	appendFunc(sl)
	/*
		 - 当是切片（slice）时，表达式 s[low : high] 中的 high，最大的取值范围对应着切片的容量（cap），不是单纯的长度（len）。
		   因此，fmt.Println(sl[:10]) 可以输出容量范围内的值，并不会出现越界。
		- fmt.Println(sl) 因为该切片 len 值为 0，没有指定最大索引值，high 则取 len 值，导致输出结果为空。
	*/
	fmt.Println(sl) // 等价于 	fmt.Println(sl[:])
	fmt.Println(sl[:10])
}
