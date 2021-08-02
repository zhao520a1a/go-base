/*
将输入的字符串解析为整数切片；

所有自定义包实现者应该遵守的最佳实践：
1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic ()
2）向包的调用者返回错误值（而不是 panic）。

当没有东西需要转换或者转换成整数失败时，这个包会 panic   --- （在函数 fields2numbers 中）
然后可导出的 Parse 函数会从 panic 中 recover 并用所有这些信息返回一个错误给调用者。
*/

package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	Index int
	Word  string
	Err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("error parsing %q as int", e.Word)
}

func Parse(intput string) (numbers []int, err error) {
	//在包内部，总是应该从 panic 中 recover
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fileds := strings.Fields(intput)
	numbers = fileds2numbers(fileds)
	return
}

func fileds2numbers(fileds []string) (numbers []int) {
	if len(fileds) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fileds {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
