/*
结构体的定义
*/

package main

import (
	"fmt"
	"strings"
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

func main() {
	//pers1是结构体类型变量
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	pers1.int = 1
	pers1.telephone = "10086"
	pers1.email = "123@qq.com"
	upPerson(&pers1)
	fmt.Println(pers1)

	//pers2 和 pers3 是一个结构体类型变量的指针
	var pers2 = new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	pers2.int = 2
	upPerson(pers2)
	fmt.Println(pers2)

	pers3 := &Person{"Chris", "Woodward", 3, contact{"10010", "hello@163.com"}}
	pers4 := Person{"Chris", "Woodward", 3, contact{"10010", "hello@163.com"}}
	fmt.Println(pers3)
	fmt.Println(pers4)

	perMap := make(map[int]*Person)
	perMap[pers3.int] = pers3
	fmt.Println("===", perMap[0])

}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
