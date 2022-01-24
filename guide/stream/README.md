### 流处理

Stream 能让我们支持链式调用和函数编程的风格来实现数据的处理，看起来数据像是在流水线一样不断的实时流转加工，最终被汇总。Stream 的实现思想就是将数据处理流程抽象成了一个数据流，每次加工后返回一个新的流供使用。

### 设计实现

#### 如何实现链式调用

链式调用，创建对象用到的 builder 模式可以达到链式调用效果。实际上 Stream 实现类似链式的效果原理也是一样的，每次调用完后都创建一个新的 Stream 返回给用户。

```
// 去除重复item
Distinct(keyFunc KeyFunc) Stream
// 按条件过滤item
Filter(filterFunc FilterFunc, opts ...Option) Stream
```

#### 如何实现流水线的处理效果

所谓的流水线可以理解为数据在 Stream 中的存储容器，在 go 中我们可以使用 channel 作为数据的管道，达到 Stream 链式调用执行多个操作时异步非阻塞效果。

#### 如何支持并行处理

数据加工本质上是在处理 channel 中的数据，那么要实现并行处理无非是并行消费 channel 而已，利用 goroutine 协程、WaitGroup 机制可以非常方便的实现并行处理。

核心逻辑是利用 channel 当做管道，数据当做水流，不断的用协程接收/写入数据到 channel 中达到异步非阻塞的效果，很难想象在 go 中 300 多行的代码就能实现如此强大的组件。

实现高效的基础来源三个语言特性：

- channel
- 协程
- 函数式编程

### [源码分析](https://mp.weixin.qq.com/s/t3INtSfFSmv-nsJqLmdPew)

![img.png](img.png)

### API 使用

Stream 的工作流程其实也属于生产消费者模型，整个流程跟工厂中的生产流程非常相似，尝试先定义一下 Stream 的生命周期：

创建阶段/数据获取（原料） 加工阶段/中间处理（流水线加工） 汇总阶段/终结操作（最终产品） 下面围绕 stream 的三个生命周期开始定义 API：

#### 创建阶段 API

##### channel 创建 Range

``` go
  return Stream{  
    source: source,  
  }  
}
```

##### 可变参数模式创建 Just

```
// 通过可变参数模式创建 stream，channel 写完后及时 close 是个好习惯。
func Just(items ...interface{}) Stream {
  source := make(chan interface{}, len(items))
  for _, item := range items {
    source <- item
  }
  close(source)
  return Range(source)
}
```

##### 函数创建 From

```
// 通过函数创建 Stream
func From(generate GenerateFunc) Stream {
  source := make(chan interface{})
  threading.GoSafe(func() {
    defer close(source)
    generate(source)
  })
  return Range(source)
}
```

#### 加工阶段 API

``` 
  // 去除重复item
  Distinct(keyFunc KeyFunc) Stream
  // 按条件过滤item
  Filter(filterFunc FilterFunc, opts ...Option) Stream
  // 分组
  Group(fn KeyFunc) Stream
  // 返回前n个元素
  Head(n int64) Stream
  // 返回后n个元素
  Tail(n int64) Stream
  // 获取第一个元素
  First() interface{}
  // 获取最后一个元素
  Last() interface{}
  // 转换对象
  Map(fn MapFunc, opts ...Option) Stream
  // 合并item到slice生成新的stream
  Merge() Stream
  // 反转
  Reverse() Stream
  // 排序
  Sort(fn LessFunc) Stream
  // 作用在每个item上
  Walk(fn WalkFunc, opts ...Option) Stream
  // 聚合其他Stream
  Concat(streams ...Stream) Stream
``` 

#### 汇总阶段 API

``` 
  // 检查是否全部匹配
  AllMatch(fn PredicateFunc) bool
  // 检查是否存在至少一项匹配
  AnyMatch(fn PredicateFunc) bool
  // 检查全部不匹配
  NoneMatch(fn PredicateFunc) bool
  // 统计数量
  Count() int
  // 清空stream
  Done()
  // 对所有元素执行操作
  ForAll(fn ForAllFunc)
  // 对每个元素执行操作
  ForEach(fn ForEachFunc)


```

