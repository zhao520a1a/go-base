/*
拼接字符串，比较标准库中的几种方法，看下哪种性能更好？
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	sss = "hello world!"
	cnt = 10000
)

var expected = strings.Repeat(sss, 10000)

// +
func BenchmarkStringConcat(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < cnt; i++ {
			str += sss
		}
		result = str
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

//  fmt.Sprintf()
func BenchmarkStringSprintf(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < cnt; i++ {
			str = fmt.Sprintf("%s%s", str, sss)
		}
		result = str
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

// strings.Join()
func BenchmarkStringJoin(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var str string
		for i := 0; i < cnt; i++ {
			str = strings.Join([]string{str, sss}, "")
		}
		result = str
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

// bytes.Buffer.WriteString()
func BenchmarkBufferWrite(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		for i := 0; i < cnt; i++ {
			buf.WriteString(sss)
		}
		result = buf.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

// append()
func BenchmarkBytesAppend(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var bbb []byte

		for i := 0; i < cnt; i++ {
			bbb = append(bbb, sss...)
		}
		result = string(bbb)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

// strings.Builder.WriteString()
func BenchmarkStringBuilder(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		var builder strings.Builder

		for i := 0; i < cnt; i++ {
			builder.WriteString(sss)
		}

		result = builder.String()
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

// copy()
func BenchmarkStringCopy(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		tsl := len(sss) * cnt
		bs := make([]byte, tsl)
		bl := 0

		for bl < tsl {
			bl += copy(bs[bl:], sss)
		}

		result = string(bs)
		// 		result = convert.Bytes2str(bs)  性能再提升
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}
