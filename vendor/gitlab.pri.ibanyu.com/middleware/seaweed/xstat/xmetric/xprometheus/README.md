## 简介
`xstat/xmetric/xprometheus`是打点统计基础库，支持Counter、Gauge、Histogram等类型，数据上报方式为`pull`模式，业务方可以只简单的定义自己的metric，然后计数，其余的一切都不需要管。如果是脚本类程序需要打点，可以选用`push`模式。

`xstat/xmetric/sys`为golang服务运行时信息统计，包含进程内存、进程CPU、gc暂停时间、goroutine数量等等，同时通过`xprometheus`将对应信息上报到prometheus。

除非有个性化需求，大多数的打点数据都会通过埋点方式集成到框架、基础库内。如果是自己使用，建议大家可以将指标统一定义在包内的metrics.go文件内，然后在包内统一计数，例如[sys](../../sys/metrics.go)，
## 使用说明
### Counter类型
- 简介

`Counter`类型是一个只增不减(系统重置除外)的计数器，常见监控指标: http_requests_total、container_cpu_usage_seconds_total、exception_total等
- 示例
```golang
const(
    namespace = "requests_http"
    subsystem = "count"
)
var(
// grafana配置metric的名称为: Namespace+"_"+Subsystem+"_"+Name，也可以只是用name。
// With 为顺序的、带标签的k-v对，可以动态新增标签
_metricRequest = NewCounter(&CounterVecOpts{
		Namespace:  namespace,
		Subsystem:  subsystem,
		Name:       "total",
		Help:       "This is the help string.",
		LabelNames: []string{"service", "ip"},
    })
)
// Inc固定增加1
_metricRequest.With("service", "picturebook", "ip", "192.168.1.1").Inc(
// Add增加对应数字
_metricRequest.With("service", "picturebook", "ip", "192.168.1.1").Add(float64(10))
    .
    .
)
```
### Gauge类型
- 简介

`Gauge`是一个可增可减的仪表盘，常见监控指标: 当前内存、当前cpu使用率、当前连接数等

- 示例

```golang
const(
    namespace = "runtime_resource"
)
var(
    // cpu usage
	_metricCpuUsage = xprometheus.NewGauge(&xprometheus.GaugeVecOpts{
		Namespace:  namespace,
		Subsystem:  "cpu_usage",
		Name:       "current",
		Help:       "cup usage percentage",
		LabelNames: []string{"service", "instance"},
    })
)
_metricCpuUsage.With("service", "picturebook", "instance", "server.id.1").Set(float64(66))
        .
        .
```
### Histogram类型
- 简介

`Histogram`类型用于统计和分析样本的分布情况，常见监控指标: 接口响应时间
- 示例

```golang
const(
    namespace = "requests_http"
    subsystem = "duration"
)
buckets := []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500}
var(
    	_metricRequestDuration = NewHistogram(&HistogramVecOpts{
		Namespace:  namespace,
		Subsystem:  subsystem,
		Name:       "current",
		Help:       "This is the help string for the histogram.",
		LabelNames: []string{"service", "instance"},
		Buckets:    buckets,
	})
)
_metricRequestDuration.With("service", "picturebook", "instance", "serve.id.1").Observe(float64(16.8))
```
## 打点监听
1.如果是基于`roc框架`自动生成的工程，框架层面会自动注册，不需要显示注册监听。

2.如果没有基于`roc框架`，通过注册`http.Handle("/metrics", promhttp.Handler())`即可。
## 打点地址发现
1.目前roc框架会通过注册metric监听信息到etcd，实现自动发现，所以对于框架生成的服务，直接进行打点即可，不需要关心监听和发现逻辑。

2.如果没有基于`roc框架`，可以让运维单独进行地址配置。
## Q&A
1.在调用某个构造函数，如NewHistogram(...)其中的LabelNames是必须携带的吗？    
LabelNames是prometheus注册阶段需要标识的标签，是`必须要携带的`，即使With里也带了完整的标签对，但是这个标签对更对的是为了标签赋值不出错，不能起到注册作用，如果两者不一致，目前的行为是panic(推动程序必须测试过才上线)

2.roc框架什么版本支持自动监听、打点地址发现?    
目前v1.2.17及以上版本支持自动监听、打点地址发现。

3.如何起名字比较合适?    
建议Namespace用组名、Subsystem用服务名、Name为你定义的有意义的点的名字，打点库会自动组成namespace_subsystem_name的全名，这样不可能出现名字冲突

4.有没有打点的最佳实践呢？   
比如需不需要每次都New一个新的对象
目前建议如果点比较多，使用一个单独的`metrics.go`文件，定义打点变量，类似于[sys](../../sys/metrics.go),然后在具体打点位置直接使用变量，`一定不要每次都new一个对象，只需要new一次`

5.什么情况下会panic?    
5.1目前使用打点库，roc框架至少需要`>=v1.2.17`，否则会panic(因为没有注册)   
5.2重复new新的对象打点会panic(duplicate metrics collector registration attempted)    
5.3new时所传`LabelNames`和`With(..`的标签不匹配