package benchmark_test

import (
	"encoding/json"
	"os"
	"runtime"
	"runtime/debug"
	"testing"
	"time"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

var (
	debugSyncGC  = os.Getenv("JSON_SYNC_GC") != ""
	debugAsyncGC = os.Getenv("JSON_NO_ASYNC_GC") == ""
)

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

var _GenericValue interface{}
var _BindingValue TwitterStruct

func init() {
	_ = json.Unmarshal([]byte(TwitterJson), &_GenericValue)
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
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
