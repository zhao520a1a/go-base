package sort

import (
	"fmt"
	"testing"
)

func TestSortMap(t *testing.T) {
	testMap := make(map[string]int)
	testMap["a"] = 1
	testMap["d"] = 4
	testMap["e"] = 5
	testMap["b"] = 2
	testMap["c"] = 3

	pairList := sortMapByValue(testMap)
	fmt.Printf("pairList : %v", pairList)
}
