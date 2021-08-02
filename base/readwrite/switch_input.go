/*
从键盘读取输入，使用了 switch 语句：
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	//注意：Unix 和 Windows 的行结束符是不同的！
	// For Unix: test-demo with delimiter "\n", for Windows: test-demo with "\r\n"
	switch input {
	case "Philip\n", "Ivo\n":
		fmt.Println("Welcome ", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
