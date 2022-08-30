/*
用函数可以作为（其它函数的）参数的事实来使用高阶函数
*/
package fun

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	ford := &Car{"Fiesta", "福特", 2008}
	bmw := &Car{"XL 450", "宝马", 2011}
	merc := &Car{"D600", "梅赛德斯", 2009}
	bmw2 := &Car{"X 800", "宝马", 2008}
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	//查找
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Facturers == "宝马") && (car.BuildYear > 2010)
	})
	fmt.Println("全部数据: ", allCars)
	fmt.Println("查找结果: ", allNewBMWs)

	//获取其他类型的集合
	facturerArr := allCars.Map(func(car *Car) Any {
		return car.Facturers
	})

	//将[]Any转换为[]string
	facturers := make([]string, 0)
	for _, v := range facturerArr {
		if s, ok := v.(string); ok {
			facturers = append(facturers, s)
		}
	}
	fmt.Println("全部厂家: ", facturerArr)

	//按车辆厂家对Cars集合做归类,返回Map集合
	//facturers := []string{"福特", "梅赛德斯", "路虎", "宝马", "奥迪"}
	sortedAppender, sortedCars := MakeSortedAppender(facturers)
	allCars.Process(sortedAppender)
	BMWCount := len(sortedCars["宝马"])
	fmt.Println("最终结果: ", sortedCars)
	fmt.Println("宝马厂家的车辆数量：", BMWCount)
}
