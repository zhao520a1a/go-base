## Benchmarks

### 测试数据

小[Small](../testdata/small.go) ：400B，11 key，深度 3 层； 中[Medium](../testdata/medium.go)：13KB，300+ key，深度 4
层（实际业务数据，其中有大量的嵌套 JSON string)； 大[Large](../testdata/large.json)：635KB，10000+ key，深度 6 层。