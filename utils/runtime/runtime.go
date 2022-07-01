package runtime

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetFunInfo 获取当前文件名和函数名
// Go 1.7+ 建议使用 runtime.CallersFrames 而不是 runtime.FuncForPC ;
func GetFunInfo() (string, string) {
	var fileName, funName string
	if pc, file, _, ok := runtime.Caller(1); ok {
		// get caller uppercase package name
		filePathItems := strings.Split(file, "/")
		fileName = filePathItems[len(filePathItems)-1]

		funFullName := runtime.FuncForPC(pc).Name()
		// note: Exclude anonymous functions
		funPathItems := strings.Split(strings.TrimRight(funFullName, ".func1"), ".")
		funName = funPathItems[len(funPathItems)-1]

	}
	return fileName, funName
}

// GetFunName 获取当前函数名
func GetFunName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func GetTrace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

// GoID 获取协程ID
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
