# go-base
学习Go过程中，总结记录一些 Go 的语法练习Demo、实用的工具包
工具包： https://github.com/zhao520a1a/go-utils


### 好玩的测试
- Go 字符串拼接的 7 种姿势，哪个最优秀？[https://shockerli.net/post/golang-concat-string/]

- Go by Example 中文：字符串格式化？
[https://books.studygolang.com/gobyexample/string-formatting/]


### 交叉编译
``` ~~~~~~~~
GOOS=linux GOARCH=amd64 go build  hello.go

nohup ./hello >> nohup.log 2>&1
``` 
 
### 其他
[Go mod 代理](https://goproxy.io/zh/)
