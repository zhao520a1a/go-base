# go-base
学习Go过程中，总结记录一些 Go 的语法练习Demo、实用的工具包

工具包：gitlab.pri.ibanyu.com/quality/dry.git 




### 好玩的测试
- Go 字符串拼接的 7 种姿势，哪个最优秀？[https://shockerli.net/post/golang-concat-string/]

- Go by Example 中文：字符串格式化？
[https://books.studygolang.com/gobyexample/string-formatting/]



### 启动服务
```
./service.git --serv=base/uauth --stype=local --logdir=./log --skey=30 &

```

### 交叉编译
``` ~~~~~~~~
GOOS=linux GOARCH=amd64 go build  hello.go

nohup ./hello >> nohup.log 2>&1
```

### 规范
#### Http请求Url格式
> 整体的api路径起名规范：  xxxapi/{服务分组}/{服务名}/{api 路径}

#### 异常处理
> 异常封装，层层上报
> 不建议在controller、dao中直接记录错误日志，可以将要记录的错误内容封装到返回的Error中，统一有最上层方法打印，避免一个错误重复多次！
使用：https://zhenghe-md.github.io/blog/2020/10/05/Go-Error-Handling-Research/

### 其他
[Go mod 代理](https://goproxy.io/zh/)