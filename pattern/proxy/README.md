## 代理模式（Proxy Pattern）

## 适用场景：
代理模式 (Proxy Pattern)，可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。

#### 声明

``` go
package proxy

import (
	"fmt"
)

type Seller interface {
	sell(name string)
}

// StationProxy 代理了 Station:代理类中持有被代理类对象，并且和被代理类对象实现了同一接口。
type StationProxy struct {
	station *Station
}

func (m *StationProxy) sell(name string) {
	if m.station.stock > 0 {
		m.station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, m.station.stock)
	} else {
		fmt.Println("票已售空")
	}
}

type Station struct {
	stock int // 库存
}

func (m *Station) sell(name string) {
	if m.stock > 0 {
		m.stock--
		fmt.Printf("火车站：%s买了一张票,剩余：%d \n", name, m.stock)
	} else {
		fmt.Println("票已售空")
	}
}
```
