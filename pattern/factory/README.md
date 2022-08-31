## 工厂模式（Factory Pattern）
工厂模式（Factory Pattern）是面向对象编程中的常用模式。可以通过使用多种不同的工厂模式，来使代码更简洁明了。

## 适用场景：
因为单例模式保证了实例的全局唯一性，而且只被初始化一次，所以比较适合全局共享一个实例，且只需要被初始化一次的场景，例如数据库实例、全局配置、全局任务池等。

### 应用模式
- 简单工厂模式
-	抽象工厂模式:抽象工厂模式和简单工厂模式的唯一区别，就是它返回的是接口而不是结构体。这样可以在不公开内部实现的情况下，让调用者使用你提供的各种功能。
-	工厂方法模式：在简单工厂模式中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，就要向工厂中传入一个参数，获取对应的对象；如果要增加一种产品，就要在工厂中修改创建产品的函数。这会导致耦合性过高，这时我们就可以使用工厂方法模式。

#### 声明
注意：在实际开发中，我建议返回**非指针的实例**，因为我们主要是想通过创建实例，调用其提供的方法，而不是对实例做更改。如果需要对实例做更改，可以实现 SetXXX 的方法。通过返回非指针的实例，可以确保实例的属性，避免属性被意外 / 任意修改。

``` go
package factory

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
```
 
