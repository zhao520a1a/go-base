

``` shell
➜  hello go test -bench=. -benchtime=3s -run=none
// Benchmark 名字 - CPU     循环次数          平均每次执行时间 
BenchmarkSprintf-8      50000000               109 ns/op
PASS
//  哪个目录下执行go test         累计耗时
ok      flysnow.org/hello       5.628s

```

### 参考资料
https://my.oschina.net/solate/blog/3034188

https://zhuanlan.zhihu.com/p/80578541