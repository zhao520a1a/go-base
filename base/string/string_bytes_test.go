package string

import (
	"strings"
	"testing"
	"unsafe"
)

/*
字符串与字节数组转换，根本原因是对底层字节数组的复制。
参考:https://segmentfault.com/a/1190000005006351
*/

var s = strings.Repeat("a", 1024)

func convert1() {
	b := []byte(s)
	_ = string(b)
}

func convert2() {
	b := str2bytes(s)
	_ = bytes2str(b)
}

// str2bytes  从 ptype 输出的结构来看，string 可看做 [2]uintptr，而 [ ]byte 则是 [3]uintptr，这便于我们编写代码，无需额外定义结构类型。如此，str2bytes 只需构建 [3]uintptr{ptr, len, len}，而 bytes2str 更简单，直接转换指针类型，忽略掉 cap 即可。
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 测试 string => byte
func BenchmarkStr2bytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStr2bytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str2bytes(s)
	}
}

// 测试 string 和 byte 互转
func BenchmarkConvert1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert1()
	}
}

func BenchmarkConvert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert2()
	}
}
