package method

import (
	"fmt"
	"strings"
	"testing"
)

type Person struct {
	firstName, lastName string
	int                 //匿名字段/内嵌字段
	contact
}

type contact struct {
	telephone string
	email     string
}

func TestCallMethod(t *testing.T) {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	pers1.int = 1
	pers1.telephone = "10086"
	pers1.email = "123@qq.com"
	fmt.Printf("原始struct地址是：%p\n", &pers1)

	//调用函数传参
	upPerson1(&pers1)
	upPerson2(pers1)

	//调用方法:指针调用方法
	var pers2 *Person
	pers2 = &pers1
	fmt.Printf("原始struct地址是：%p\n", pers2)
	pers2.fullName("-")

	//调用方法：值调用方法
	pers1.SetFirstName("Eric")
	fmt.Printf("per2.firstName = %s \n", pers2.firstName)

	//调用内嵌对象的方法
	pers1.changeTelephone("10000")
	fmt.Println(pers1)
	fmt.Println(pers1.String())

}

//引用传递
func upPerson1(p *Person) {
	fmt.Printf("upPerson1 函数里接收到struct的内存地址是：%p\n", p)
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

//值传递
func upPerson2(p Person) {
	fmt.Printf("upPerson2 函数里接收到struct的内存地址是：%p\n", &p)
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

//无论调用者是值还是指针，方法都支持运行: 将值对象作为接收者,值传递；
func (p Person) fullName(segment string) string {
	fmt.Printf("fullName 函数里接收到struct的内存地址是：%p\n", &p)
	return p.firstName + segment + p.lastName
}

//无论调用者是值还是指针，方法都支持运行,Go会帮我们自动解引用: 将指针对象作为接收者,引用传递；
func (p *Person) SetFirstName(newName string) {
	fmt.Printf("SetFirstName 函数里接收到struct的内存地址是：%p\n", p)
	p.firstName = newName
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (c *contact) changeTelephone(newTeleNum string) {
	fmt.Println("匿名类型的同名方法")
	c.telephone = newTeleNum
}

func (p *Person) changeTelephone(newTeleNum string) {
	fmt.Println("覆盖匿名类型的同名方法")
	p.telephone = newTeleNum
}

func (p *Person) String() string {
	return "(" + p.firstName + "-" + p.lastName + ")"
}
