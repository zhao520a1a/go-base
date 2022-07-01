package example

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterCache(t *testing.T) {
	var wg sync.WaitGroup
	var gNum = 100000
	var kLen = 100
	manager := NewCounterMapManager()
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < kLen; j++ {
				manager.IncrBy(strconv.Itoa(j), 1)
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, kLen, manager.dataMap.Count())
	for _, k := range manager.dataMap.Keys() {
		v, _ := manager.dataMap.Get(k)
		val, ok := v.(*int64)
		if !ok {
			continue
		}
		if *val != int64(gNum) {
			t.Errorf("key %s value %v \n", k, v)
			return
		}
	}
}
