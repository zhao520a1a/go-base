package slicedemo

import (
	"fmt"
	"testing"
)

func TestCreateArray(t *testing.T) {
	/*
		创建数组
	*/
	// 创建数组方式1
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	fmt.Println("arr1:", arr1)

	// 数组是一种 值类型（不像 C/C++ 中是指向首元素的指针）,数组间赋值时，需要内存拷贝,这样两个数组就有了不同内存空间；
	arr2 := arr1
	arr2[0] = -1
	fmt.Println("arr2", arr2)

	// 创建数组方式2: arr3是指针类型
	arr3 := new([5]int)
	arr3[0] = 88
	arr3[1] = 99
	arr3[2] = 66
	fmt.Println("arr3", arr3)

	arr2 = *arr3
	arr2[1] = 100
	fmt.Println("arr2", arr2)

	// 数组常量
	arr4 := [5]string{"a", "b", "c", "d"}
	arr5 := [...]string{"a", "b", "c", "d"}
	arr6 := []string{"a", "b", "c", "d"}
	arr7 := [...]string{3: "Chris", 4: "Ron"}
	fmt.Println("arr4", arr4)
	fmt.Println("arr5", arr5)
	fmt.Println("arr6", arr6)
	fmt.Println("arr7", arr7)

	// 多维数组
	var screen [2][3]int
	for x := 0; x < 2; x++ {
		for y := 0; y < 3; y++ {
			screen[x][y] = 1
		}
	}
	fmt.Println("多维数组 :", screen)

	fmt.Println("求和：", Sum(&arr1))

}

// 传递数组的指针给函数
func Sum(arr *[5]int) (sum int) {
	for _, v := range arr { // 使用*arr返回数组是没有必要的
		sum += v
	}
	return
}
