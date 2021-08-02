/*
从键盘和标准输入 os.Stdin 读取输入，最简单的办法是使用 fmt 包提供的 Scan 和 Sscan 开头的函数。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  string
	format                 = "%f / %d / %s"
)

var inputReader *bufio.Reader

func main() {
	fmt.Println("Please enter your full name: ")
	//扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Println("Hi %s %s!\n", firstName, lastName)

	//使用 bufio 包提供的缓冲读取
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n') //  输入：56.12 / 5212 / Go
	if err == nil {
		fmt.Println("The input was: %s", input)
	}

	//从字符串读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
