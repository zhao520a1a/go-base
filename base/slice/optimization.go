package slicedemo

import (
	"io/ioutil"
	"regexp"
)

/*
切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。
当没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存
*/

//功能：将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片。
var digitRegexp = regexp.MustCompile("[0-9]+")

//返回的 b []byte 指向的底层是整个文件的数据。只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存。
func FindDigits1(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

//可以通过拷贝我们需要的部分到一个新的切片中
func FindDigits2(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
