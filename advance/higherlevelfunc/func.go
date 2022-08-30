package fun

import "strconv"

type Any interface {
}

type Car struct {
	Mode      string //车型
	Facturers string //厂商
	BuildYear int    //生产年份
	//...
}

type Cars []*Car

func (c Car) String() string {
	return strconv.Itoa(c.BuildYear) + "-" + c.Facturers + "-" + c.Mode
}

// Process 功能：通用函数，执行入参函数；  -- 函数入参
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// FindAll 功能：将满足条件的找出来，添加到集合中。 -- 将有返回值的函数入参
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// Map 功能： 通过cars集合获取其他类型结合 -- 将外层函数作为变量传入高级函数中
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result = append(result, f(c))
		ix++
	})
	return result
}

// MakeSortedAppender 功能：产生特定的添加函数(根据不同的厂商添加汽车到不同的集合)和 map集合 (等价于Java中Map<String,List<Car>>集合)
func MakeSortedAppender(features []string) (func(car *Car), map[string]Cars) {
	//初始化carMap及其key值，用于后续使用
	carMap := make(map[string]Cars)
	carMap["其他"] = make([]*Car, 0)
	for _, v := range features {
		carMap[v] = make([]*Car, 0)
	}

	//添加函数
	appender := func(c *Car) {
		if _, ok := carMap[c.Facturers]; ok {
			carMap[c.Facturers] = append(carMap[c.Facturers], c)
		} else {
			carMap["其他"] = append(carMap["其他"], c)
		}
	}
	return appender, carMap
}
