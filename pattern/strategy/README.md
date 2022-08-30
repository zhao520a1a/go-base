## 策略模式（Strategy Pattern）

策略模式定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。

## 适用场景：

经常要根据不同的场景，采取不同的措施，也就是不同的策略；避免通过 if ... else ... 的形式来调用不同的计算方式

### 应用模式

#### 声明

``` go
// 定义一个策略类
type IStrategy interface {
	do(int, int) int
}

// 策略实现：加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 策略实现：减
type sub struct{}

func (*sub) do(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (m *Operator) setStrategy(strategy IStrategy) {
	m.strategy = strategy
}

// 调用策略中的方法
func (m *Operator) calculate(a, b int) int {
	return m.strategy.do(a, b)
}

```

#### 使用

``` go
 // 可以随意更换策略，而不影响 Operator 的所有实现。
func TestStrategy(t *testing.T) {
	operator := Operator{}
	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(&sub{})
	result = operator.calculate(2, 1)
	fmt.Println("sub:", result)
}
```

### 参考

