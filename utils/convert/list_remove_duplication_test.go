package convert

import (
	"fmt"
	"sort"
	"testing"
)

/*
去除slice和list的重复元素，非常有用的功能
*/

func TestDuplicate(t *testing.T) {
	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	fmt.Println(Duplicate(b))

	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8}
	sort.Ints(c)
	fmt.Println(Duplicate(c))
}
