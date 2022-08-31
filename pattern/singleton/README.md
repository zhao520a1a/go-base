## 模板模式（Singleton Pattern）
单例模式指的是全局只有一个实例，并且它负责创建自己的对象。单例模式不仅有利于减少内存开支，还有减少系统性能开销、防止多个实例产生冲突等优点。

## 适用场景：
因为单例模式保证了实例的全局唯一性，而且只被初始化一次，所以比较适合全局共享一个实例，且只需要被初始化一次的场景，例如数据库实例、全局配置、全局任务池等。

### 应用模式
单例模式又分为饿汉方式和懒汉方式。
饿汉方式指全局的单例实例在包被加载时创建，而懒汉方式指全局的单例实例在第一次被使用时创建。你可以看到，这种命名方式非常形象地体现了它们不同的特点。

#### 声明

``` go

package singleton

import "sync"

var ins *singleton
var once sync.Once

type singleton struct {
}

func GetSingleton() *singleton {
  // 使用 once.Do 可以确保 ins 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
```
 
