package pattern

import "fmt"

type Person interface {
	Greet()
}

type Coder struct {
	name string
	age  int
}

func (m Coder) Greet() {
	fmt.Printf("Hi, My name is %s", m.name)
}

// 简单工厂模式
func NewCoder1(name string, age int) Coder {
	return Coder{
		name: name,
		age:  age,
	}
}

// 抽象工厂模式
func NewCoder2(name string, age int) Person {
	return Coder{
		name: name,
		age:  age,
	}
}

/*
工厂方法模式: 创建具有默认年龄的工厂
	f := NewCoderFactory(30)
	f("golden")
    f("mike")
*/
func NewCoderFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Coder{
			name: name,
			age:  age,
		}
	}
}
