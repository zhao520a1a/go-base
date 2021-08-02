package main

import "fmt"

var (
	map1 = map[int]string{
		2: "韩梅梅",
		3: "李雷",
		1: "小明同学"}
)

/*
键值对调
*/
func main() {

	invMap := make(map[string]int, len(map1))
	for k, v := range map1 {
		invMap[v] = k
	}
	fmt.Printf("%v \n", map1)
	fmt.Printf("%v \n", invMap)

}
