/*
指针可以调用值方法！
- 类型 *T 的可调用方法集 = 接受者为 *T +  T 的所有方法集
- 类型 T 的可调用方法集 = 接受者为 T 的所有方法，不包含接受者为 *T 的方法
*/

package method

import (
	"fmt"
	"testing"
)

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func TestPrtCall(t *testing.T) {
	sVals := map[int]S{1: {"old"}}
	sVals[1].Read()
	//sVals[1].Write("new") 不能编译通过，
	fmt.Println(sVals[1].Read())

	sPrts := map[int]*S{1: {"old"}}
	sPrts[1].Read() //指针调用值接受器方法
	sPrts[1].Write("new")
	fmt.Println(sPrts[1].Read())
}
