package main

import "fmt"

// 定义一个接口泛型约束
type Addable[T any] interface {
	Add(T) T
}

// 定义一个泛型类型
type Pair[T int | float64] struct {
	first, second T
}

func (p *Pair[T]) Add(value T) T {
	return p.first + value
}

// 定义一个泛型函数
func Sum[T ~int | ~float64](a, b T) T {
	return a + b
}

func main() {
	// 使用泛型类型
	pair := Pair[int]{1, 2}
	fmt.Println("Pair:", pair)
	fmt.Println("Pair.Add():", pair.Add(1))

	// 使用泛型函数
	sum := Sum(1, 2) // 整数类型
	fmt.Println("Sum:", sum)

	sumFloat := Sum(1.5, 2.3) // 浮点类型
	fmt.Println("SumFloat:", sumFloat)
}
