/*
slices 和 maps 包含了指向底层数据的指针，注意用户对暴露内部状态的 map 或 slice 的修改。
*/

package convert

import "sync"

type Stats struct {
	mu sync.Mutex

	counters map[string]int
	data     []string
}

//内部字段赋值时要复制，防止外部修改data内容后影响 s.data
func (s Stats) SetData(data []string) {
	s.data = make([]string, len(data))
	copy(s.data, data)
}

func (s Stats) getSnapshotCounts() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}
