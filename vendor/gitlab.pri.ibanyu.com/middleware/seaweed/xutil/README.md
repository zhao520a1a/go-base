## 简介
xutil主要包含各种工具函数，方便基础库、业务进行一些公共操作。

### Future模块
为了更加安全的使用goroutine，提供了future功能，内部自带recover，使用示例如下:

```go
func Foo(ctx context.Context) (string, error) {
    return "bar", nil
}

ctx := context.Background()
var rv string
f := xutil.NewFuture(func() error {
    var err error
    rv, err = Foo(ctx)
    return err
})
err := f.Get(ctx)

fmt.Println(rv)
```