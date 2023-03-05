package benchmark_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkEncodeStream_Jsoniter(b *testing.B) {
	var o = map[string]interface{}{
		"a": `<` + strings.Repeat("1", 1024) + `>`,
		"b": json.RawMessage(` [ ` + strings.Repeat(" ", 1024) + ` ] `),
	}

	b.Run("single", func(b *testing.B) {
		var w = bytes.NewBuffer(nil)
		var jt = jsoniter.Config{
			ValidateJsonRawMessage: true,
		}.Froze()
		var enc = jt.NewEncoder(w)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(o)
			w.Reset()
		}
	})

	b.Run("double", func(b *testing.B) {
		var w = bytes.NewBuffer(nil)
		var jt = jsoniter.Config{
			ValidateJsonRawMessage: true,
		}.Froze()
		var enc = jt.NewEncoder(w)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(o)
			_ = enc.Encode(o)
			w.Reset()
		}
	})

	b.Run("compatible", func(b *testing.B) {
		var w = bytes.NewBuffer(nil)
		var jt = jsoniter.Config{
			ValidateJsonRawMessage: true,
			EscapeHTML:             true,
			SortMapKeys:            true,
		}.Froze()
		var enc = jt.NewEncoder(w)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = enc.Encode(o)
			w.Reset()
		}
	})
}
