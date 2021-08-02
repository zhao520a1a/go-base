/*
返回获取结构体中标签
*/
package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "我是bool类型字段"
	field2 int    "我是int类型字段"
	field3 string "我是string类型字段"
}

func main() {
	tt := TagType{true, 1, "Golden"}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField.Tag)
}
