package main

import (
	"fmt"
	"sync"
)

// Cache 是一个泛型缓存结构，可以存储任意类型的键值对
type Cache[K comparable, V any] struct {
	data  map[K]V
	mutex sync.Mutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]V)}
}

// Set 向缓存中添加或更新一个键值对
func (c *Cache[K, V]) Set(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

// Get 从缓存中获取一个键对应的值
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {
	cache := NewCache[int, string]()
	cache.Set(1, "apple")
	cache.Set(2, "banana")

	value, ok := cache.Get(1)
	if ok {
		fmt.Println("Cached value:", value)
	} else {
		fmt.Println("Value not found")
	}
}
