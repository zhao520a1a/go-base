/*
反射结构体
*/
package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	s1, s2, s3, S4 string
}

var t1 interface{} = T1{"a", "b", "c", "d"}
var t2 = T1{"a", "b", "c", "d"}

func main() {
	fmt.Println(reflect.TypeOf(t1))
	value := reflect.ValueOf(t1)
	fmt.Println(value.Kind())
	fmt.Println(value)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(i, "-", value.Field(i))
	}
	fmt.Println("--", value.Method(0).Call(nil))

	elem := reflect.ValueOf(&t2).Elem()
	//结构体中只有被导出字段（首字母大写）才是可设置的； 否则运行时会报错： using value obtained using unexported field
	//elem.Field(0).SetString("1")
	elem.Field(3).SetString("ff")
	fmt.Println(t2)

	//打印结构体中的字段信息（注意：不能打印非导出字段）
	//for i := 0; i < elem.NumField(); i++ {
	//	f := elem.Field(i)
	//	fmt.Printf("%d: %s %s = %v \n",i,elem.Type().Field(i).DepartmentName,f.Type(),f.Interface())
	//}

}

func (n T1) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3 + " - " + n.S4
}
