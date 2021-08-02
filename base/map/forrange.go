package main

import "fmt"

/*
遍历map
注意 map-demo 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
*/
func main() {

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0

	for key, value := range map1 {
		fmt.Printf("key: %d - value:%f \n", key, value)
	}

	for key := range map1 {
		fmt.Printf("key: %d - value:%f \n", key, map1[key])
	}

}
