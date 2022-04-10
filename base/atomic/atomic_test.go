package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

// 让 AddUint32 做减法
func TestAddUint32(t *testing.T) {
	diff := -2

	// 方式一
	a := uint32(2)
	b := atomic.AddUint32(&a, uint32(int32(diff)))
	fmt.Printf("a = %d, b = %d \n", a, b)

	// 方式二：
	c := uint32(2)
	d := atomic.AddUint32(&c, ^uint32(-diff-1))
	fmt.Printf("c = %d, d = %d \n", c, d)
}
