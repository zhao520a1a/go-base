/*
使用工厂方法创建结构体实例
*/
package main

import (
	"fmt"
	"unsafe"
)

//结构体名称为小写，根据可见性规则是该类型变成私有的，则可以禁止其他包使用new函数创建，从而强制使用工厂方法
type file struct {
	fd   int    //文件描述符
	name string //文件名
}

func main() {
	a := &file{10, "./test-demo.txt"}
	f := NewFile(10, "./test-demo.txt")
	fmt.Println(a)
	fmt.Println(f)

	//结构体类型 T 的一个实例占用了多少内存
	fmt.Println(unsafe.Sizeof(file{}))
}

//工厂方法
func NewFile(fd int, name string) *file {
	if fd < 0 {
		return nil
	}
	return &file{fd, name}
}
