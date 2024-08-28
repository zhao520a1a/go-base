package main

import (
	"fmt"
	"time"
)

// 定义泛型栈
type Stack[T any] struct {
	elements []T
}

// Push方法
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop方法
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // 类型T的零值
		return zero, false
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, true
}

func main() {
	// 实例化一个泛型类型的栈
	stack := Stack[any]{}
	stack.Push(1)
	stack.Push("Hello")
	stack.Push(time.Now())

	// 依次弹出栈中的元素
	for {
		element, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println(element)
	}
}
