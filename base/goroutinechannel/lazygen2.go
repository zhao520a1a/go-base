/*
 通过巧妙地使用空接口、闭包和高阶函数，实现一个通用的惰性生产器的工厂函数 BuildLazyEvaluator。
*/

package main

import "fmt"

type Any interface{}
type EvalFunc func(Any) (Any, Any)

/*
输入: 一个函数(需要计算出下一个返回值以及下一个状态参数)和一个初始状态
返回: 一个无参、返回值是生成序列的函数
*/
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	//创建一个通道和无限循环的 go 协程,将返回值被放到了该通道中，
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	//返回函数从该通道中取得该返回值
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}
