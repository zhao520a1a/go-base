/*
使用选择依赖注入方式避免使用可变的全局变量
*/
package main

import (
	"fmt"
	"testing"
	"time"
)

var _timeNow = time.Now

type signer struct {
	now func() time.Time
}

//使用选择依赖注入方式避免改变全局变量,这样既适用于函数指针又适用于其他值类型
func newSigner() *signer {
	return &signer{
		now: time.Now,
	}
}

func (s *signer) Sign() string {
	//now := _timeNow()  要尽量避免可变全局变量
	now := s.now()
	return now.String()
}

func TestSigner(t *testing.T) {
	s := newSigner()
	s.now = func() time.Time {
		return time.Now()
	}

	fmt.Sprintln(s.Sign())
}
