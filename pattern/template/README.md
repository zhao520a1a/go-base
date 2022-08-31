## 模板模式（Template Pattern）
模版模式 (Template Pattern) 定义一个操作中算法的骨架，将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方。


## 适用场景：
将一些步骤延迟到子类中。这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。

### 应用模式

#### 声明

``` go
package template

import (
	"fmt"
	"testing"
)

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似与一个抽象类
type CookMenu struct {
}

func (*CookMenu) fire() {
	fmt.Println("开火🔥")
}

func (*CookMenu) cooke() {

}

func (*CookMenu) outfire() {
	fmt.Println("关火🧯")
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cook() {
	fmt.Println("做西红柿🍅")
}

type ChaoJiDan struct {
	CookMenu
}

func (*ChaoJiDan) cook() {
	fmt.Println("炒一个鸡蛋🥚")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}
```

#### 使用

``` go
func TestTemplate(t *testing.T) {
	xiHongShi := &XiHongShi{}
	doCook(xiHongShi)
	chaoJiDan := &ChaoJiDan{}
	doCook(chaoJiDan)
}
```
