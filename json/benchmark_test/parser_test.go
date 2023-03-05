package benchmark_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

/*
- 测试数据：中等大小 JSON 字符串
- 测试场景：unmarshal 为 interface{}
- 测试操作：
*/

func BenchmarkParser_Gjson(b *testing.B) {
	gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
		if !value.Exists() {
			b.Fatal(value.Index)
		}
		_ = value.Value()
		return true
	})
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
			if !value.Exists() {
				b.Fatal(value.Index)
			}
			_ = value.Value()
			return true
		})
	}
}

func BenchmarkParser_Jsoniter(b *testing.B) {
	v := jsoniter.Get([]byte(TwitterJson)).GetInterface()
	if v == nil {
		b.Fatal(v)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = jsoniter.Get([]byte(TwitterJson)).GetInterface()
	}
}

func BenchmarkParser_Parallel_Gjson(b *testing.B) {
	gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
		if !value.Exists() {
			b.Fatal(value.Index)
		}
		return true
	})
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gjson.Parse(TwitterJson).ForEach(func(key, value gjson.Result) bool {
				if !value.Exists() {
					b.Fatal(value.Index)
				}
				_ = value.Value()
				return true
			})
		}
	})
}

func BenchmarkParser_Parallel_Jsoniter(b *testing.B) {
	var bv = []byte(TwitterJson)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var out interface{}
			_ = jsoniter.Unmarshal(bv, &out)
		}
	})
}
