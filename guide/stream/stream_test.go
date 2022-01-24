package stream

import (
	"testing"

	. "github.com/tal-tech/go-zero/core/fx"
)

func TestDistinct(t *testing.T) {
	// 1 2 3 4 5
	Just(1, 2, 3, 3, 4, 5, 5).Distinct(func(item interface{}) interface{} {
		return item
	}).ForEach(func(item interface{}) {
		t.Log(item)
	})

	// 1 2 3 4
	Just(1, 2, 3, 3, 4, 5, 5).Distinct(func(item interface{}) interface{} {
		uid := item.(int)
		// 对大于4的item进行特殊去重逻辑,最终只保留一个>3的item
		if uid > 3 {
			return 4
		}
		return item
	}).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 过滤
func TestInternalStream_Filter(t *testing.T) {
	// 保留偶数 2,4
	Just(1, 2, 3, 4, 5).Filter(func(item interface{}) bool {
		return item.(int)%2 == 0
	}).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 遍历执行：因为内部采用了协程机制异步执行读取和写入数据所以新的 Stream 中 channel 里面的数据顺序是随机的。
func Test_Stream_Walk(t *testing.T) {
	// 返回 300,100,200
	Just(1, 2, 3).Walk(func(item interface{}, pip chan<- interface{}) {
		pip <- item.(int) * 100
	}, WithWorkers(3)).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 一比一映射转换
func TestInternalStream_Map(t *testing.T) {
	Just(1, 2, 3, 4, 5, 2, 2, 2, 2, 2, 2).Map(func(item interface{}) interface{} {
		return item.(int) * 10
	}).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 分组 Group
func TestInternalStream_Group(t *testing.T) {
	var groups [][]int
	Just(10, 11, 20, 21).Group(func(item interface{}) interface{} {
		v := item.(int)
		return v / 10
	}).ForEach(func(item interface{}) {
		t.Log(item)
		v := item.([]interface{})
		var group []int
		for _, each := range v {
			group = append(group, each.(int))
		}
		groups = append(groups, group)
	})
	t.Log(groups)

}

// 获取前 n 个元素 Head
// 返回1,2
func TestInternalStream_Head(t *testing.T) {
	Just(1, 2, 3, 4, 5).Head(2).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

func TestInternalStream_Tail(t *testing.T) {
	// 4,5
	Just(1, 2, 3, 4, 5).Tail(2).ForEach(func(item interface{}) {
		t.Log(item)
	})

	// 1,2,3,4,5
	Just(1, 2, 3, 4, 5).Tail(6).ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 反转 Reverse
func TestInternalStream_Reverse(t *testing.T) {
	Just(1, 2, 3, 4, 5).Reverse().ForEach(func(item interface{}) {
		t.Log(item)
	})
}

// 排序
// 5,4,3,2,1
func TestInternalStream_Sort(t *testing.T) {
	Just(1, 2, 3, 4, 5).Sort(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	}).ForEach(func(item interface{}) {
		t.Log(item)
	})
}
