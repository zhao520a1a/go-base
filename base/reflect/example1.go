package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	//var x = "22"
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("type:", reflect.TypeOf(x).Kind())

	v := reflect.ValueOf(x) //值拷贝
	fmt.Println("value:", v.Kind())
	fmt.Println("value:", v.Type())
	fmt.Println("value:", v)
	//fmt.Println("value:",v.Float())
	fmt.Println("value:", v.Interface())
	fmt.Println("value:", v.Interface().(float64))
	fmt.Println("value can Set?:", v.CanSet())

	//设置/修改值
	v = reflect.ValueOf(&x) //引用拷贝
	fmt.Println("value can Set?:", v.CanSet())
	v = v.Elem()
	fmt.Println("elem-value:", v)
	fmt.Println("value can Set?:", v.CanSet())
	v.SetFloat(123)
	fmt.Println(v.Interface())
	fmt.Println(x)

}
