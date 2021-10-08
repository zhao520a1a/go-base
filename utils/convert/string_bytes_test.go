package convert

// 字符串与字节数组转换，根本原因是对底层字节数组的复制。

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func convert1() {
	b := []byte(s)
	_ = string(b)
}

func convert2() {
	b := str2bytes(s)
	_ = bytes2str(b)
}

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
