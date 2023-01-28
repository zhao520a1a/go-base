package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

/*
在 Go 中，字符串是以 UTF-8 为格式进行存储的
如果字符串中出现中文字符不能直接调用 len 函数来统计字符串字符长度，
，在字符串上调用 len 函数，取得的是字符串包含的 byte 的个数。
*/
func TestStringEncode(t *testing.T) {
	str := "Hello, 世界"
	// 获取字符串包含的 字符 的个数。
	l1 := len([]rune(str))
	l2 := utf8.RuneCountInString(str)
	// 判断字节nil在字符串s中出现的次数，没有找到则返回-1，如果为空字符串("")则返回字符串的长度+1
	l3 := bytes.Count([]byte(str), nil) - 1
	// 统计字符""出现的次数
	l4 := strings.Count(str, "") - 1
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)

	// 获取字符串包含的 byte 的个数。
	l5 := len(str)
	fmt.Println(l5)

	// 遍历字符串时，迭代变量下标不连续
	for i, c := range str {
		fmt.Printf("%d:%q[%x] \n", i, c, []byte(string(c)))
	}

}

func TestTrimLeft(t *testing.T) {
	fmt.Println(strings.TrimLeft("121212121222113123", "12"))
	fmt.Println(strings.TrimPrefix("12123123", "12"))
}
