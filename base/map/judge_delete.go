package main

import (
	"fmt"
)

/*
判断键值对是否存在 & 删除元素
*/
func main() {

	map1 := make(map[string]string)
	map1["北京"] = "中国的首都"
	map1["平壤"] = "朝鲜的首都"
	map1["首尔"] = "韩国的首都"

	/*
		判断键值对是否存在
	*/
	//写法1
	value, isPresent := map1["上海"]
	if isPresent {
		fmt.Println("该键不存在集合中,值为：" + value)
	} else {
		fmt.Println("该键不存在集合中！")
	}

	//写法2
	if _, isPresent = map1["北京"]; isPresent {
		fmt.Println("该键存在集合中！")
		//...
	}

	/*
		删除元素：如果 key1 不存在，该操作不会产生错误。
	*/
	delete(map1, "北京")
	_, isPresent = map1["北京"]
	fmt.Println("该键是否存在集合中：", isPresent)

}
