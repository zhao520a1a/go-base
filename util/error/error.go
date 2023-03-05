/*

Go 中有多种声明错误（Error) 的选项：
errors.New 对于简单静态字符串的错误
fmt.Errorf 用于格式化的错误字符串
实现 Error() 方法的自定义类型
用 "pkg/errors".Wrap 的 Wrapped errors

错误具有以下传播方式
如果不需要额外信息的简单错误,则使用errors.New 足够了。
如果客户需要检测并处理此错误，则使用自定义类型并实现该 Error() 方法。
如果正在传播下游函数返回的错误，则使用错误包装Error Wrapping 以便错误消息提供更多上下文 ,errors.Cause 可用于提取原始错误。
如果调用者不需要检测或处理的特定错误情况等其他情况，则使用 fmt.Errorf。

*/

package error

import (
	"errors"
	"fmt"
)

// 不需要额外信息的简单错误
var ErrCouldNotOpen = errors.New("could not open1")

// 生成ID失败
type errIdGenRpc struct {
	Fun string
	Msg error
}

// 参数无效
type errParamInvalid struct {
	Fun string
	Msg string
}

// 包装错误
type errWarping struct {
	Fun string
	Msg string
	Err error
}

// 没有找到
type errNotFound struct {
	Fun string
	Msg string
}

// 数据已存在
type errAlreadyExist struct {
	Fun string
	Msg string
}

var _ error = (*errIdGenRpc)(nil)
var _ error = (*errParamInvalid)(nil)
var _ error = (*errWarping)(nil)
var _ error = (*errAlreadyExist)(nil)
var _ error = (*errNotFound)(nil)
var _ error = (*errWarping)(nil)

func (e errIdGenRpc) Error() string {
	return fmt.Sprintf("%s --> IdgenRPC generate DepartmentId error: %s ", e.Fun, e.Msg.Error())
}

func (e errParamInvalid) Error() string {
	return fmt.Sprintf("%s --> invalid parameter: %s ", e.Fun, e.Msg)
}

func (e errWarping) Error() string {
	return fmt.Sprintf("%s --> %s error:%s ", e.Fun, e.Msg, e.Err)
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("%s --> %s not found! ", e.Fun, e.Msg)
}

func (e errAlreadyExist) Error() string {
	return fmt.Sprintf("%s --> %s already exist", e.Fun, e.Msg)
}

func GetStrContent(value ...interface{}) string {
	if len(value) == 0 {
		return ""
	}
	return fmt.Sprintln("params：", value)
}

func NewInvalidParamErr(fun, msg string) error {
	return &errParamInvalid{Fun: fun, Msg: msg}
}

func IsParamInvalidError(err error) bool {
	_, ok := err.(*errParamInvalid)
	return ok
}

func NewIdGenRpcErr(fun string, msg error) error {
	return &errIdGenRpc{Fun: fun, Msg: msg}
}

func IsIdGenRpcError(err error) bool {
	_, ok := err.(*errIdGenRpc)
	return ok
}

func NewWrappingErr(fun, msg string, err error) error {
	return &errWarping{Fun: fun, Msg: msg, Err: err}
}

func NewWrappingError(fun string, err error) error {
	return &errWarping{Fun: fun, Err: err}
}

func IsWrappingError(err error) bool {
	_, ok := err.(*errWarping)
	return ok
}

func NewNotFoundErr(fun, msg string) error {
	return &errNotFound{Fun: fun, Msg: msg}
}

// IsNotFoundError 最好公开匹配器功能以检查错误
func IsNotFoundError(err error) bool {
	_, ok := err.(*errNotFound)
	return ok
}

func NewAlreadyExisError(fun, msg string) error {
	return &errAlreadyExist{Fun: fun, Msg: msg}
}

func IsAlreadyExisError(err error) bool {
	_, ok := err.(*errAlreadyExist)
	return ok
}
