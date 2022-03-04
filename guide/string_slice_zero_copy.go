package main

import (
	"unsafe"
)

/*
使用zero copy的方式处理slice和string，来节省内存
*/
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

/*
	从上可以看出来slice 和stringStruct就差一个cap，补齐就好
*/
type stringStruct struct {
	str unsafe.Pointer
	len int
}

// Bytes2str 这里利用的事unsafe.Pointer指针互转，把byte直接转换成string了
func Bytes2str(b []byte) string {
	if b != nil && len(b) > 0 {
		return *(*string)(unsafe.Pointer(&b))
	}
	return ""
}

// Str2bytesV1 这里转换其实有个问题，unsafe.Pointer引用的对象如果中途被修改，结果也是相应会变化，造成内容不一致。
func Str2bytesV1(s string) []byte {
	if len(s) > 0 {
		x := (*[2]uintptr)(unsafe.Pointer(&s))
		h := [3]uintptr{x[0], x[1], x[1]}
		return *(*[]byte)(unsafe.Pointer(&h))
	}
	return []byte{}
}

// Str2bytesV2 正确的使用方式，通过构造新的结构，指向原有的地址，解释外层变量变化，也不会影响到内层，代价就是多一个指针操作
func Str2bytesV2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
