package util

import (
	"context"
	"fmt"
)

// 将输入数据拆分成分页获取
func ListByPage(ctx context.Context) {
	uids := make([]int64, 100)
	pageSize := 50
	for i := 0; i < len(uids); i += pageSize {
		uidList := uids[i:MinInt(i+pageSize, len(uids))]
		//rpc call with uidList
		fmt.Println(uidList)
	}
	return
}

// 分页返回
func PageQuery(offset, limit int64) {

	var processList []int
	total := int64(len(processList))
	if total > offset {
		if total > offset+limit {
			processList = processList[offset : offset+limit]
		} else {
			processList = processList[offset:total]
		}
	}

	processList = processList[MinInt64(offset, total):MinInt64(offset+limit, total)]

}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
