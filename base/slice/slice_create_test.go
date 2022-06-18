package slicedemo

import (
	"fmt"
	"testing"
)

/*
切片是一个 长度可变的数组。
*/
func TestSlice(t *testing.T) {
	var arr1 [6]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println("数组：", arr1)

	slice1 := arr1[2:4]
	fmt.Println("初始化：", slice1)

	slice1 = slice1[0:3]
	//slice1 = slice1[-1:3] -- 重新分片只能向后移动，不允许向前启动，否则会导致编译错误。
	fmt.Println("重新分片：", slice1)

	//切片可以反复扩展直到占据整个相关数组。
	slice1 = slice1[:cap(slice1)]
	fmt.Println("切片扩展到上限：", slice1)
	fmt.Println("传递切片给函数:", sum(arr1[:]))

	//用make创建切片
	slice2 := make([]int, 10, 50)
	for j := 0; j < len(slice2); j++ {
		slice2[j] = j
	}
	fmt.Println("make创建切片:", slice2)

	//切片常量
	slice3 := []string{"春", "夏", "秋", "冬"}
	slice4 := [...]string{"春", "夏", "秋", "冬"}
	fmt.Println(slice3)
	fmt.Println(slice4)

	//错误做法- 用new来创建切片 此时：*p == nil，切片的len和cap都为0
	var p *[]int = new([]int)
	fmt.Println("new 创建切片", *p)

}

//传递切片给函数
func sum(arr []int) int {
	s := 0
	for _, value := range arr {
		s += value
	}
	return s
}

/*
创建map类型的切片
	1.分配切片
	2.分配切片中每个元素

*/
func TestCreateMapSlice(t *testing.T) {

	items := make([]map[int]string, 5)

	//通过索引使用切片的 map-demo 元素
	for i := range items {
		items[i] = make(map[int]string)
		items[i][1] = "a"
	}
	fmt.Printf("items: %v \n", items)

	//错误的方式： item只是map的一个拷贝, 实质切片中的map没有得到初始化。
	for _, item := range items {
		item = make(map[int]string)
		item[1] = "b"
	}
	fmt.Printf("items: %v \n", items)

}

/*
切片的复制和追加
*/
func TestCopyAndAppend(t *testing.T) {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	//复制
	n := copy(slTo, slFrom)
	fmt.Println("复制的元素数量：", n)
	fmt.Println(cap(slTo), "-", slTo)

	//追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。返回的切片可能已经指向一个不同的相关数组了。
	slTo = append(slTo, 4, 5, 6, 7)
	fmt.Println(cap(slTo), "-", slTo)

	slTo = AppendInt(slTo, 8, 9)
	fmt.Println(cap(slTo), "-", slTo)

}

// 手动实现append方法
func AppendInt(slice []int, elems ...int) []int {
	m := len(slice)
	n := m + len(elems)
	if n > cap(slice) {
		newSlice := make([]int, (n+1)*2) // 创建新的分片
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], elems)
	return slice
}
