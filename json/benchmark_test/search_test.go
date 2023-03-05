package benchmark_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/bytedance/sonic/ast"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/zhao520a1a/go-base/json/testdata"
)

var TwitterJson = testdata.TwitterJson

/*
- 测试数据：中等大小 JSON 字符串
- 测试场景：查找（get）& 修改（set）
- 测试操作：指定某种规则的查找路径（一般是 key 与 index 的集合），获取需要的那部分 JSON value 并处理。
*/

/* --- 获取一个Json数组中某个元素的字段值 ---*/
func BenchmarkGetOne_Gjson(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ast := gjson.Get(TwitterJson, "statuses.3.id")
		node := ast.Int()
		if node != 249279667666817024 {
			b.Fail()
		}
	}
}

func BenchmarkGetOne_Jsoniter(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	data := []byte(TwitterJson)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ast := jsoniter.Get(data, "statuses", 3, "id")
		node := ast.ToInt()
		if node != 249279667666817024 {
			b.Fail()
		}
	}
}

func BenchmarkGetOne_Sonic(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	ast := ast.NewSearcher(TwitterJson)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node, err := ast.GetByPath("statuses", 3, "id")
		if err != nil {
			b.Fatal(err)
		}
		x, _ := node.Int64()
		if x != 249279667666817024 {
			b.Fatal(node.Interface())
		}
	}
}

func BenchmarkGetOne_Parallel_Gjson(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ast := gjson.Get(TwitterJson, "statuses.3.id")
			node := ast.Int()
			if node != 249279667666817024 {
				b.Fail()
			}
		}
	})
}

func BenchmarkGetOne_Parallel_Jsoniter(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ast := jsoniter.Get([]byte(TwitterJson), "statuses", 3, "id")
			node := ast.ToInt()
			if node != 249279667666817024 {
				b.Fail()
			}
		}
	})
}

func BenchmarkGetOne_Parallel_Sonic(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ast := ast.NewSearcher(TwitterJson)
			node, err := ast.GetByPath("statuses", 3, "id")
			if err != nil {
				b.Fatal(err)
			}
			x, _ := node.Int64()
			if x != 249279667666817024 {
				b.Fatal(node.Interface())
			}
		}
	})
}

/*
连续 7 次获取字段
*/
//func BenchmarkGetSeven_Gjson(b *testing.B) {
//	b.SetBytes(int64(len(TwitterJson)))
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		ast := gjson.Parse(TwitterJson)
//		node := ast.Get("statuses.3.id")
//		node = ast.Get("statuses.3.user.entities.description")
//		node = ast.Get("statuses.3.user.entities.url.urls")
//		node = ast.Get("statuses.3.user.entities.url")
//		node = ast.Get("statuses.3.user.created_at")
//		node = ast.Get("statuses.3.user.name")
//		node = ast.Get("statuses.3.text")
//		if node.Value() == nil {
//			b.Fail()
//		}
//	}
//}
//
//func BenchmarkGetSeven_Jsoniter(b *testing.B) {
//	b.SetBytes(int64(len(TwitterJson)))
//	b.ResetTimer()
//	data := []byte(TwitterJson)
//	for i := 0; i < b.N; i++ {
//		ast := jsoniter.Get(data)
//		node := ast.Get("statuses", 3, "id")
//		node = ast.Get("statuses", 3, "user", "entities", "description")
//		node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
//		node = ast.Get("statuses", 3, "user", "entities", "url")
//		node = ast.Get("statuses", 3, "user", "created_at")
//		node = ast.Get("statuses", 3, "user", "name")
//		node = ast.Get("statuses", 3, "text")
//		if node.LastError() != nil {
//			b.Fail()
//		}
//	}
//}
//
//func BenchmarkGetSeven_Parallel_Gjson(b *testing.B) {
//	b.SetBytes(int64(len(TwitterJson)))
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			ast := gjson.Parse(TwitterJson)
//			node := ast.Get("statuses.3.id")
//			node = ast.Get("statuses.3.user.entities.description")
//			node = ast.Get("statuses.3.user.entities.url.urls")
//			node = ast.Get("statuses.3.user.entities.url")
//			node = ast.Get("statuses.3.user.created_at")
//			node = ast.Get("statuses.3.user.name")
//			node = ast.Get("statuses.3.text")
//			if node.Value() == nil {
//				b.Fail()
//			}
//		}
//	})
//}
//
//func BenchmarkGetSeven_Parallel_Jsoniter(b *testing.B) {
//	b.SetBytes(int64(len(TwitterJson)))
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			data := []byte(TwitterJson)
//			ast := jsoniter.Get(data)
//			node := ast.Get("statuses", 3, "id")
//			node = ast.Get("statuses", 3, "user", "entities", "description")
//			node = ast.Get("statuses", 3, "user", "entities", "url", "urls")
//			node = ast.Get("statuses", 3, "user", "entities", "url")
//			node = ast.Get("statuses", 3, "user", "created_at")
//			node = ast.Get("statuses", 3, "user", "name")
//			node = ast.Get("statuses", 3, "text")
//			if node.LastError() != nil {
//				b.Fail()
//			}
//		}
//	})
//}
//

/* --- 获取一个Json对象的字段值 ---*/
func BenchmarkGetByKeys_Sonic(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	ast := ast.NewSearcher(TwitterJson)
	const _count = 4
	for i := 0; i < b.N; i++ {
		node, err := ast.GetByPath("search_metadata", "count")
		if err != nil {
			b.Fatal(err)
		}
		x, _ := node.Int64()
		if x != _count {
			b.Fatal(node.Interface())
		}
	}
}

func BenchmarkGetByKeys_JsonParser(b *testing.B) {
	b.SetBytes(int64(len(TwitterJson)))
	data := []byte(TwitterJson)
	const _count = 4
	for i := 0; i < b.N; i++ {
		value, err := jsonparser.GetInt(data, "search_metadata", "count")
		if err != nil {
			b.Fatal(err)
		}
		if value != _count {
			b.Fatal(value)
		}
	}
}

/* --- 设置一个Json数组中某个元素的字段值 ---*/
func BenchmarkSetOne_Sjson(b *testing.B) {
	path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
	_, err := sjson.Set(TwitterJson, path, math.MaxInt32)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sjson.Set(TwitterJson, path, math.MaxInt32)
	}
}

func BenchmarkSetOne_Jsoniter(b *testing.B) {
	data := []byte(TwitterJson)
	node, ok := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
	if !ok {
		b.Fatal(node)
	}

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node, _ := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
		node["id"] = math.MaxInt32
	}
}

func BenchmarkSetOne_Parallel_Sjson(b *testing.B) {
	path := fmt.Sprintf("%s.%d.%s", "statuses", 3, "id")
	_, err := sjson.Set(TwitterJson, path, math.MaxInt32)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sjson.Set(TwitterJson, path, math.MaxInt32)
		}
	})
}

func BenchmarkSetOne_Parallel_Jsoniter(b *testing.B) {
	data := []byte(TwitterJson)
	node, ok := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
	if !ok {
		b.Fatal(node)
	}

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			node, _ := jsoniter.Get(data, "statuses", 3).GetInterface().(map[string]interface{})
			node["id"] = math.MaxInt32
		}
	})
}
