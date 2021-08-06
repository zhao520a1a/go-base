package main

import (
	"fmt"
	"testing"
)

// 并不确定interface类型时候，也可以使用 anyInterface.(type) 结合 switch case 来做判断

type User struct{ Name string }

func TestTypeAssertion(t *testing.T) {
	any := User{Name: "fidding"}
	test(any)

	userList := []User{{
		Name: "Tom",
	}, {
		Name: "Jerry",
	}}
	test(userList)

	testMap := map[string]string{
		"Name": "牛郎",
		"Sex":  "男",
	}
	test(testMap)

}

func test(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	case User: // 可以看到op即为将interface转为User struct类型，并使用其Name对象
		op, ok := value.(User)
		fmt.Println(op.Name, ok)
	case []User:
		if users, ok := value.([]User); ok {
			for _, user := range users {
				fmt.Println(user.Name)
			}
		}
	case map[string]string:
		if op, ok := value.(map[string]string); ok {
			fmt.Println(op["Name"])
			fmt.Println(op["Sex"])
		}

	default:
		fmt.Println("unknown")
	}
}
