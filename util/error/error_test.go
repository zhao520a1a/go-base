package error

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	file := "testfile.txt"

	if err := open1(file); err != nil {
		if err == ErrCouldNotOpen {
			//handle
			fmt.Println(err)
		} else {
			panic("unkonwn error")

		}
	}

	if err := open2(file); err != nil {
		if IsNotFoundError(err) {
			//handle
			fmt.Println(err)
		} else {
			panic("unkonwn error")
		}
	}

	if err := open3(file); err != nil {
		//handle
		fmt.Println(err)
	}

}

func open1(file string) error {
	return ErrCouldNotOpen
}

func open2(file string) error {
	fun := "open2 -->"
	return NewNotFoundErr(fun, file)
}

//Error Wrapping

func open3(file string) error {
	fun := "open3 -->"
	//请避免使用“failed to”之类的短语以保持上下文简洁，这些短语会陈述明显的内容，但是，一旦将错误发送到另一个系统，就应该明确消息是错误消息（例如使用err标记，或在日志中以”Failed”为前缀）。
	err := open1(file)
	if err != nil {
		return NewWrappingError(fun, err)
	}
	return nil

}
