## 简介
xconfig是的内部配置中心库，支持apollo、etcd等配置中心数据的读取，支持配置热加载。

## 使用说明

### 构造 ConfigCenter 对象 

```golang
// 构造ConfigCenter对象，入口参数为配置类型，类型直接使用xconfig.ConfigerTypeApollo
configCenter, err := xconfig.NewConfigCenter(apollo.ConfigTypeApollo)
if err != nil {
	fmt.Printf("new configer center err:%s", err.Error())
	return
}
```

### 监听 ConfigCenter 配置变更
***注：改步骤可以省略，如果不需要动态感知变更事件***

1.构造监听对象，入参是 事件处理函数

2.注册监听对象

```golang
observer := xconfig.NewConfigObserver(func(ctx context.Context, event *xconfig.ChangeEvent) {
   	// Todo 处理变更事件
})
configCenter.RegisterObserver(ctx, observer)
```

### 读取配置

```golang
configCenter.GetBool(ctx, "key_bool")
configCenter.GetBoolWithNamespace(ctx, "test", "key_bool_2")
```


### 代码样例

* 可运行示例

[xconfig/example/main.go](xconfig/example/main.go)

* 代码示例

```golang
package main

import (
	"context"
	"fmt"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xconfig/apollo"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xconfig"
)

func main() {
	ctx := context.TODO()
	//Step1 : 构造实例
	configCenter, err := xconfig.NewConfigCenter(ctx, apollo.ConfigTypeApollo, "base/test", []string{"application", "test"})
	if err != nil {
		fmt.Printf("new configer center err:%s", err.Error())
		return
	}
	
	//Step2 : 监听变更（改步骤可以省略，如果不需要动态感知变更事件）
	observer := xconfig.NewConfigObserver(func(ctx context.Context, event *xconfig.ChangeEvent) {
		// Todo 处理变更事件
	})
	configCenter.RegisterObserver(ctx, observer)

	//Step3 : 获取配置Value
	configCenter.GetBool(ctx, "key_bool")
	configCenter.GetBoolWithNamespace(ctx, "test", "key_bool_2")
}
```
