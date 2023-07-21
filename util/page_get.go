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
		// rpc call with uidList
		fmt.Println(uidList)
	}
	return
}

// 分页返回
func PageQuery(dataList []string, offset, limit int) (pageDataList []string) {
	total := len(dataList)
	if total > offset {
		if total > offset+limit {
			pageDataList = dataList[offset : offset+limit]
			return
		}
		pageDataList = dataList[offset:total]
		return
	}
	pageDataList = dataList[MinInt(offset, total):MinInt(offset+limit, total)]
	return
}

// 可以改写成泛型模式
func Page[Data any](dataList []*Data, offset, limit int) (pageDataList []*Data) {
	// sort.SliceStable(dataList, func(i, j int) bool {
	// 	return dataList[i].Date > dataList[j].Date
	// })
	total := len(dataList)
	if total > offset {
		if total > offset+limit {
			pageDataList = dataList[offset : offset+limit]
			return
		}
		pageDataList = dataList[offset:total]
		return
	}
	pageDataList = dataList[MinInt(offset, total):MinInt(offset+limit, total)]
	return
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
