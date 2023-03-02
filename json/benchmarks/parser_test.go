package benchmark_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

/*
- 测试数据：中等大小 JSON 字符串
- 测试场景：解析
- 测试操作：指定某种规则的查找路径（一般是 key 与 index 的集合），获取需要的那部分 JSON value 并处理。
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

func BenchmarkParseOne_Gjson(b *testing.B) {
	ast := gjson.Parse(TwitterJson)
	node := ast.Get("statuses.2.id")
	v := node.Int()
	if v != 249289491129438208 {
		b.Fatal(node)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ast := gjson.Parse(TwitterJson)
		node := ast.Get("statuses.2.id")
		_ = node.Int()
	}
}

func BenchmarkParseOne_Jsoniter(b *testing.B) {
	data := []byte(TwitterJson)
	ast := jsoniter.Get(data, "statuses", 2, "id")
	node := ast.ToInt()
	if node != 249289491129438208 {
		b.Fail()
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ast := jsoniter.Get(data, "statuses", 2, "id")
		_ = ast.ToInt()
	}
}

func BenchmarkParseOne_Parallel_Gjson(b *testing.B) {
	ast := gjson.Parse(TwitterJson)
	node := ast.Get("statuses.2.id")
	v := node.Int()
	if v != 249289491129438208 {
		b.Fatal(node)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ast := gjson.Parse(TwitterJson)
			node := ast.Get("statuses.2.id")
			_ = node.Int()
		}
	})
}

func BenchmarkParseOne_Parallel_Jsoniter(b *testing.B) {
	data := []byte(TwitterJson)
	ast := jsoniter.Get(data, "statuses", 2, "id")
	node := ast.ToInt()
	if node != 249289491129438208 {
		b.Fail()
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := []byte(TwitterJson)
			ast := jsoniter.Get(data, "statuses", 2, "id")
			_ = ast.ToInt()
		}
	})
}

/*
获取 7 次
*/

func BenchmarkParseSeven_Gjson(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ast := gjson.Parse(TwitterJson)
		node := ast.Get("statuses.3.id")
		node = ast.Get("statuses.3.user.entities.description")
		node = ast.Get("statuses.3.user.entities.url.urls")
		node = ast.Get("statuses.3.user.entities.url")
		node = ast.Get("statuses.3.user.created_at")
		node = ast.Get("statuses.3.user.name")
		node = ast.Get("statuses.3.text")
		if node.Value() == nil {
			b.Fail()
		}
	}
}

func BenchmarkParseSeven_Jsoniter(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	data := []byte(TwitterJson)
	for i := 0; i < b.N; i++ {
		ast := jsoniter.Get(data)
		node := ast.Get("statuses", 3, "id")
		node = ast.Get("statuses", 3, "user", "entities", "description")
		node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
		node = ast.Get("statuses", 3, "user", "entities", "url")
		node = ast.Get("statuses", 3, "user", "created_at")
		node = ast.Get("statuses", 3, "user", "name")
		node = ast.Get("statuses", 3, "text")
		if node.LastError() != nil {
			b.Fail()
		}
	}
}

func BenchmarkParseSeven_Parallel_Gjson(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ast := gjson.Parse(TwitterJson)
			node := ast.Get("statuses.3.id")
			node = ast.Get("statuses.3.user.entities.description")
			node = ast.Get("statuses.3.user.entities.url.urls")
			node = ast.Get("statuses.3.user.entities.url")
			node = ast.Get("statuses.3.user.created_at")
			node = ast.Get("statuses.3.user.name")
			node = ast.Get("statuses.3.text")
			if node.Value() == nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkParseSeven_Parallel_Jsoniter(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := []byte(TwitterJson)
			ast := jsoniter.Get(data)
			node := ast.Get("statuses", 3, "id")
			node = ast.Get("statuses", 3, "user", "entities", "description")
			node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
			node = ast.Get("statuses", 3, "user", "entities", "url")
			node = ast.Get("statuses", 3, "user", "created_at")
			node = ast.Get("statuses", 3, "user", "name")
			node = ast.Get("statuses", 3, "text")
			if node.LastError() != nil {
				b.Fail()
			}
		}
	})
}
