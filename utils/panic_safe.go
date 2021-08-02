package util

import (
	"fmt"
	"runtime"
)

// SafeFunWrapperWithArgs 给运行的函数f封装，避免panic导致全局退出
func SafeFunWrapperWithArgs(f func(args ...interface{}), args ...interface{}) (err error) {

	defer func() {
		if ok := recover(); ok != nil {
			err = DumpStack(ok)
		}
	}()

	f(args...)
	return
}

func SafeFunWrapper(f func()) (err error) {

	defer func() {
		if ok := recover(); ok != nil {
			err = DumpStack(ok)
		}
	}()

	f()
	return
}

func DumpStack(e interface{}) (err error) {
	if e == nil {
		return
	}
	err = fmt.Errorf("%+v", e)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		err = fmt.Errorf("%s\t %s:%d", err.Error(), file, line)
	}
	return
}
