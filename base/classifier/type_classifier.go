/*
类型判断:
一个类型分类函数，它有一个可变长度参数，可以是任意类型的数组，它会根据数组元素的实际类型执行不同的动
*/
package classifier

import "fmt"

func Classifier(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case float32, float64:
			fmt.Printf("Param #%d is a float\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		default:
			fmt.Printf("Param #%d is unkown \n", i)
		}
	}
}
