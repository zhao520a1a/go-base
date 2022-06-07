package _map

import (
	"fmt"
	"sync"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
)

func TestSyncMap(t *testing.T) {
	// 声明 scene 类型为 sync.Map 。注意，sync Map 不能使用 make 创建。
	var scene sync.Map

	// 保存键值，是以interface{}类型进行保存
	scene.Store("green", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 获取键值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历map
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ", key, value)
		// 更新map
		scene.Delete(key)
		scene.Store("golden", 111)
		return true
	})

	fmt.Println(scene.Load("golden"))
}

func TestSyncMapV2(t *testing.T) {
	testMap := cmap.New()

	// 保存键值
	testMap.Set("green", 97)
	testMap.Set("london", 100)
	testMap.Set("egypt", 200)

	// 获取键取值
	fmt.Println(testMap.Get("london"))

	// 根据键删除对应的键值对
	testMap.Remove("london")

	// 遍历+更新：map会死锁
	testMap.IterCb(func(key string, value interface{}) {
		fmt.Println("iterate: ", key, value)
		// 更新map
		testMap.Remove(key)
		testMap.Set("golden", 111)
		return
	})
}
