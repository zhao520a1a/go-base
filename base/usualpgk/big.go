package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
对于整数的高精度计算 Go 语言中提供了 big 包。
其中包含了 math 包：有用来表示大整数的 big.Int 和表示大有理数的 big.Rat 类型
这些类型可以实现任意位类型的数字，只要内存足够大。缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多。
*/
func main() {

	//大的整型数字
	im := big.NewInt(math.MaxInt64)
	ip := big.NewInt(1)
	fmt.Println(im)
	fmt.Println(ip.Mul(im, im))
	fmt.Println(ip.Div(ip, im))
	fmt.Println(ip.Add(ip, ip))
	fmt.Println(ip)

	//大有理数
	rm := big.NewRat(1111, 2222)
	rn := big.NewRat(19, 56)
	fmt.Println(rm.Mul(rm, rn))
	fmt.Println(rm)

}
