package benchmark_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/sugawarayuuta/sonnet"

	"github.com/zhao520a1a/go-base/util/rt"
)

func init() {
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func stdDecode(s string, v interface{}, copy bool) error {
	d := json.NewDecoder(bytes.NewReader([]byte(s)))
	err := d.Decode(v)
	if err != nil {
		return err
	}
	return err
}

func sonicDecode(s string, v interface{}, copy bool) error {
	d := decoder.NewDecoder(s)
	if copy {
		d.CopyString()
	}
	err := d.Decode(v)
	if err != nil {
		return err
	}
	return err
}

func BenchmarkDecoder_Generic_StdLib(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = json.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = json.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Generic_JsonIter(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = jsoniter.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = jsoniter.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Generic_GoJson(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = gojson.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = gojson.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Generic_Sonic(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = sonic.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = sonic.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Generic_Sonic_V1(b *testing.B) {
	var w interface{}
	_ = sonicDecode(TwitterJson, &w, true)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = sonicDecode(TwitterJson, &v, true)
	}
}

func BenchmarkDecoder_Generic_Sonic_Fast(b *testing.B) {
	var w interface{}
	_ = sonicDecode(TwitterJson, &w, false)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = sonicDecode(TwitterJson, &v, false)
	}
}

func BenchmarkDecoder_Generic_Sonnet(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = sonnet.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = sonnet.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Parallel_Generic_StdLib(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = json.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = json.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_JsonIter(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = jsoniter.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = jsoniter.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_GoJson(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = gojson.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = gojson.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_Sonic(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = sonic.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = sonic.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_Sonic_V1(b *testing.B) {
	var w interface{}
	_ = sonicDecode(TwitterJson, &w, true)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = sonicDecode(TwitterJson, &v, true)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_Sonic_Fast(b *testing.B) {
	var w interface{}
	_ = sonicDecode(TwitterJson, &w, false)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = sonicDecode(TwitterJson, &v, false)
		}
	})
}

func BenchmarkDecoder_Parallel_Generic_Sonnet(b *testing.B) {
	var w interface{}
	m := []byte(TwitterJson)
	_ = sonnet.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v interface{}
			_ = sonnet.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Binding_StdLib(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = json.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = json.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Binding_JsonIter(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = jsoniter.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = jsoniter.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Binding_GoJson(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = gojson.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = gojson.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Binding_Sonic(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = sonic.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = sonic.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Binding_Sonic_V1(b *testing.B) {
	var w TwitterStruct
	_ = sonicDecode(TwitterJson, &w, true)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = sonicDecode(TwitterJson, &v, true)
	}
}

func BenchmarkDecoder_Binding_Sonic_Fast(b *testing.B) {
	var w TwitterStruct
	_ = sonicDecode(TwitterJson, &w, false)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v TwitterStruct
		_ = sonicDecode(TwitterJson, &v, false)
	}
}

func BenchmarkDecoder_Binding_Sonnet(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = sonnet.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = sonnet.Unmarshal(m, &v)
	}
}

func BenchmarkDecoder_Parallel_Binding_StdLib(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = json.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = json.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_JsonIter(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = jsoniter.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = jsoniter.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_GoJson(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = gojson.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = gojson.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_Sonic(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = sonic.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = sonic.Unmarshal(m, &v)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_Sonic_V1(b *testing.B) {
	var w TwitterStruct
	_ = sonicDecode(TwitterJson, &w, true)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = sonicDecode(TwitterJson, &v, true)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_Sonic_Fast(b *testing.B) {
	var w TwitterStruct
	_ = sonicDecode(TwitterJson, &w, false)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = sonicDecode(TwitterJson, &v, false)
		}
	})
}

func BenchmarkDecoder_Parallel_Binding_Sonnet(b *testing.B) {
	var w TwitterStruct
	m := []byte(TwitterJson)
	_ = sonnet.Unmarshal(m, &w)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var v TwitterStruct
			_ = sonnet.Unmarshal(m, &v)
		}
	})
}

func BenchmarkSkip_Sonic(b *testing.B) {
	var data = rt.Str2Mem(TwitterJson)
	if ret, _ := decoder.Skip(data); ret < 0 {
		b.Fatal()
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = decoder.Skip(data)
	}
}
