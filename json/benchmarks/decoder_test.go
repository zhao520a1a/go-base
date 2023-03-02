package benchmark_test

import (
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"

	"github.com/zhao520a1a/go-base.git/json/benchmarks/testdata"
)

func init() {
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

type TwitterStruct testdata.TwitterStruct

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
