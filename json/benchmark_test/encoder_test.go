package benchmark_test

import (
	"encoding/json"
	"os"
	"runtime"
	"runtime/debug"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"

	"github.com/zhao520a1a/go-base/json/testdata"
)

/*
- 测试数据：中等大小 JSON 字符串
- 测试场景：
    泛型（generic）编码：针对 interface{} 的 marshal 操作
    定型（binding）编解码：针对固定的schema struct 的 marshal 操作
*/

var (
	debugAsyncGC = os.Getenv("JSON_NO_ASYNC_GC") == ""
)

var _GenericValue interface{}
var _BindingValue TwitterStruct

// 小JSON数据
type TwitterStruct testdata.Book

var TwitterJson = testdata.BookJson

// 中JSON数据
//type TwitterStruct testdata.TwitterStruct
// var TwitterJson = testdata.TwitterJson

func init() {
	// 大JSON数据
	//data, _ := os.ReadFile("../testdata/large.json")
	//TwitterJson = string(data)
	_ = json.Unmarshal([]byte(TwitterJson), &_GenericValue)
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func TestMain(m *testing.M) {
	go func() {
		if !debugAsyncGC {
			return
		}
		println("Begin GC looping...")
		for {
			runtime.GC()
			debug.FreeOSMemory()
		}
		println("stop GC looping!")
	}()
	time.Sleep(time.Millisecond)
	m.Run()
}

func BenchmarkEncoder_Generic_StdLib(b *testing.B) {
	_, _ = json.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(_GenericValue)
	}
}

func BenchmarkEncoder_Generic_JsonIter(b *testing.B) {
	_, _ = jsoniter.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(_GenericValue)
	}
}

func BenchmarkEncoder_Generic_GoJson(b *testing.B) {
	_, _ = gojson.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gojson.Marshal(_GenericValue)
	}
}

func BenchmarkEncoder_Generic_Sonic(b *testing.B) {
	_, _ = sonic.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(_GenericValue)
	}
}

func BenchmarkEncoder_Generic_Sonic_V1(b *testing.B) {
	_, _ = encoder.Encode(_GenericValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Encode(_GenericValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	}
}

func BenchmarkEncoder_Generic_Sonic_Fast(b *testing.B) {
	_, _ = encoder.Encode(_GenericValue, 0)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Encode(_GenericValue, 0)
	}
}

func BenchmarkEncoder_Parallel_Generic_StdLib(b *testing.B) {
	_, _ = json.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = json.Marshal(_GenericValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Generic_JsonIter(b *testing.B) {
	_, _ = jsoniter.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = jsoniter.Marshal(_GenericValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Generic_GoJson(b *testing.B) {
	_, _ = gojson.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = gojson.Marshal(_GenericValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Generic_Sonic(b *testing.B) {
	_, _ = sonic.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = sonic.Marshal(_GenericValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Generic_Sonic_V1(b *testing.B) {
	_, _ = encoder.Encode(_GenericValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = encoder.Encode(_GenericValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
		}
	})
}

func BenchmarkEncoder_Parallel_Generic_Sonic_Fast(b *testing.B) {
	_, _ = encoder.Encode(_GenericValue, encoder.NoQuoteTextMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = encoder.Encode(_GenericValue, encoder.NoQuoteTextMarshaler)
		}
	})
}

func BenchmarkEncoder_Binding_StdLib(b *testing.B) {
	_, _ = json.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&_BindingValue)
	}
}

func BenchmarkEncoder_Binding_JsonIter(b *testing.B) {
	_, _ = jsoniter.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(&_BindingValue)
	}
}

func BenchmarkEncoder_Binding_GoJson(b *testing.B) {
	_, _ = gojson.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gojson.Marshal(&_BindingValue)
	}
}

func BenchmarkEncoder_Binding_Sonic(b *testing.B) {
	_, _ = sonic.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonic.Marshal(&_BindingValue)
	}
}

func BenchmarkEncoder_Binding_Sonic_V1(b *testing.B) {
	_, _ = encoder.Encode(&_BindingValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Encode(&_BindingValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	}
}

func BenchmarkEncoder_Binding_Sonic_Fast(b *testing.B) {
	_, _ = encoder.Encode(&_BindingValue, encoder.NoQuoteTextMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Encode(&_BindingValue, encoder.NoQuoteTextMarshaler)
	}
}

func BenchmarkEncoder_Parallel_Binding_StdLib(b *testing.B) {
	_, _ = json.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = json.Marshal(&_BindingValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Binding_JsonIter(b *testing.B) {
	_, _ = jsoniter.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = jsoniter.Marshal(&_BindingValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Binding_GoJson(b *testing.B) {
	_, _ = gojson.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = gojson.Marshal(&_BindingValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Binding_Sonic(b *testing.B) {
	_, _ = sonic.Marshal(&_BindingValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = sonic.Marshal(&_BindingValue)
		}
	})
}

func BenchmarkEncoder_Parallel_Binding_Sonic_V1(b *testing.B) {
	_, _ = encoder.Encode(&_BindingValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = encoder.Encode(&_BindingValue, encoder.SortMapKeys|encoder.EscapeHTML|encoder.CompactMarshaler)
		}
	})
}

func BenchmarkEncoder_Parallel_Binding_Sonic_Fast(b *testing.B) {
	_, _ = encoder.Encode(&_BindingValue, encoder.NoQuoteTextMarshaler)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = encoder.Encode(&_BindingValue, encoder.NoQuoteTextMarshaler)
		}
	})
}
