package _defer

import (
	"fmt"
	"strconv"
	"testing"
)

/*
输出
x = c
y = d
defer: a d
*/
func TestDefer1(t *testing.T) {
	x, y := "a", "b"

	defer func(x string) {
		println("defer:", x, y) // y 闭包引用 输出延迟后的值 y = "d"
	}(x) // 匿名函数调用，传送参数x 被复制 x="a"

	x = "c"
	y = "d"
	println("x =", x, "y =", y)
}

func TestDefer2(t *testing.T) {
	x, y := "a", "b"

	if x != "a" { // 不满足条件，不会执行下面的语句
		defer func(s string) {
			println("defer:", s, y)
		}(x)
	}
	x = "c"
	y = "d"
	println("x =", x, "y =", y)
}

func TestDefer3(t *testing.T) {
	fmt.Println(testReturnError())
}

// 测试 err 抛出情况
func testReturnError() (err error) {
	// var err error
	defer func() {
		if err != nil {
			err = fmt.Errorf("defer: %v", err)
		}
	}()
	err = fmt.Errorf("I am litter error")
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi("xx")
		if err != nil {
			return err
		}
		fmt.Println(num)
	}
	return nil
}

// 多个 defer 调用采用栈顺序
func TestDefer4(t *testing.T) {
	defer func() {
		println("defer:a")
	}()
	defer func() {
		println("defer:b")
	}()
	defer func() {
		println("defer:c")
	}()
	println("finished")
}
