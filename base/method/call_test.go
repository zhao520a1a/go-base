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

	//调用函数
	upPerson1(&pers1)
	fmt.Println(pers1)
	upPerson2(pers1)
	fmt.Println(pers1)

	//调用方法1
	fullName := pers1.fullName("-")
	fmt.Println(fullName)

	//调用方法2:指针调用方法
	var pers2 *Person
	pers2 = &pers1
	fmt.Println(pers2.firstName)

	//调用方法3：值调用方法
	pers1.SetFirstName("Eric")
	fmt.Println(pers1.firstName)

	//调用内嵌对象的方法
	pers1.changeTelephone("10000")
	fmt.Println(pers1)

}

func upPerson1(p *Person) {
	fmt.Printf("函数里接收到struct的内存地址是：%p\n", p)
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func upPerson2(p Person) {
	fmt.Printf("函数里接收到struct的内存地址是：%p\n", &p)
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

//值传递
func (p Person) fullName(segment string) string {
	return p.firstName + segment + p.lastName
}

//将指针作为接收者,引用传递；Go会帮我们自动解引用，无论调用者是值还是指针，方法都支持运行
func (p *Person) SetFirstName(newName string) {
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
