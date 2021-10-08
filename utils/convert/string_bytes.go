package convert

import "unsafe"

// 参考:https://segmentfault.com/a/1190000005006351
// 从 ptype 输出的结构来看，string 可看做 [2]uintptr，而 [ ]byte 则是 [3]uintptr，这便于我们编写代码，无需额外定义结构类型。如此，str2bytes 只需构建 [3]uintptr{ptr, len, len}，而 bytes2str 更简单，直接转换指针类型，忽略掉 cap 即可。

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
