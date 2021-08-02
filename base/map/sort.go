package main

import (
	"fmt"
	"sort"
)

/*
按Key值给Map排序
将 key（或者 value）拷贝到一个切片，再对切片排序
*/

func main() {

	keys := make([]int, len(map1))
	i := 0
	for key := range map1 {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%v \n", map1[k])
	}

}
