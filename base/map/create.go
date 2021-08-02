package main

import "fmt"

/*
 Map的声明、初始化、make
*/
func main() {
	/*
		声明与初始化
	*/
	var mapInit = map[string]int{"one": 1, "two": 2}
	var mapCopy = mapInit
	mapMake := make(map[string]float32)
	mapCopy["two"] = 3
	mapMake["key1"] = 4.5
	fmt.Printf("key-one：%d \n", mapInit["two1"])
	fmt.Printf("key-ten：%f \n", mapMake["key1"])
	fmt.Printf("mapMake：%f \n", mapMake)

	/*
		不要使用 new，永远用 make 来构造 map-demo
		注意:如果错误的使用 new () 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：
	*/
	//mapNew := new(map-demo[string-demo] float32)
	//mapNew["key1"] = 4.6

	/*
	 数组、切片和结构体不能作为 key， 任意类型都能作为Map的value
	*/
	// 将函数作为map的值
	mapFunc := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(mapFunc) //返回函数地址

	//将切片作为map的值：一个key对应多个value
	//mapSlice := make()

}
