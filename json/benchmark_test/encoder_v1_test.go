package benchmark_test

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/bytedance/sonic/encoder"

	"github.com/zhao520a1a/go-base/util/rt"
)

func BenchmarkHTMLEscape_Std(b *testing.B) {
	jsonByte := []byte(TwitterJson)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	var buf []byte
	for i := 0; i < b.N; i++ {
		out := bytes.NewBuffer(make([]byte, 0, len(TwitterJson)*6/5))
		json.HTMLEscape(out, jsonByte)
		buf = out.Bytes()
	}
	_ = buf
}

func BenchmarkHTMLEscape_Sonic(b *testing.B) {
	jsonByte := []byte(TwitterJson)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	var buf []byte
	for i := 0; i < b.N; i++ {
		buf = encoder.HTMLEscape(nil, jsonByte)
	}
	_ = buf
}

func BenchmarkValidate_Std(b *testing.B) {
	var data = rt.Str2Mem(TwitterJson)
	if !json.Valid(data) {
		b.Fatal()
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Valid(data)
	}
}

func BenchmarkValidate_Sonic(b *testing.B) {
	var data = rt.Str2Mem(TwitterJson)
	ok, s := encoder.Valid(data)
	if !ok {
		b.Fatal(s)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Valid(data)
	}
}

func BenchmarkCompact_Std(b *testing.B) {
	var data = rt.Str2Mem(TwitterJson)
	var dst = bytes.NewBuffer(nil)
	if err := json.Compact(dst, data); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst.Reset()
		_ = json.Compact(dst, data)
	}
}

type f64Bench struct {
	name  string
	float float64
}

func BenchmarkEncode_Float64(b *testing.B) {
	var bench = []f64Bench{
		{"Zero", 0},
		{"ShortDecimal", 1000},
		{"Decimal", 33909},
		{"Float", 339.7784},
		{"Exp", -5.09e75},
		{"NegExp", -5.11e-95},
		{"LongExp", 1.234567890123456e-78},
		{"Big", 123456789123456789123456789},
	}
	maxUint := "18446744073709551615"
	for i := 1; i <= len(maxUint); i++ {
		name := strconv.FormatInt(int64(i), 10) + "-Digs"
		num, _ := strconv.ParseUint(string(maxUint[:i]), 10, 64)
		bench = append(bench, f64Bench{name, float64(num)})
	}
	for _, c := range bench {
		libs := []struct {
			name string
			test func(*testing.B)
		}{{
			name: "StdLib",
			test: func(b *testing.B) {
				_, _ = json.Marshal(c.float)
				for i := 0; i < b.N; i++ {
					_, _ = json.Marshal(c.float)
				}
			},
		}, {
			name: "Sonic",
			test: func(b *testing.B) {
				_, _ = encoder.Encode(c.float, 0)
				for i := 0; i < b.N; i++ {
					_, _ = encoder.Encode(c.float, 0)
				}
			},
		}}
		for _, lib := range libs {
			name := lib.name + "_" + c.name
			b.Run(name, lib.test)
		}
	}
}

type f32Bench struct {
	name  string
	float float32
}

func BenchmarkEncode_Float32(b *testing.B) {
	var bench = []f32Bench{
		{"Zero", 0},
		{"ShortDecimal", 1000},
		{"Decimal", 33909},
		{"ExactFraction", 3.375},
		{"Point", 339.7784},
		{"Exp", -5.09e25},
		{"NegExp", -5.11e-25},
		{"Shortest", 1.234567e-8},
	}

	maxUint := "18446744073709551615"
	for i := 1; i <= len(maxUint); i++ {
		name := strconv.FormatInt(int64(i), 10) + "-Digs"
		num, _ := strconv.ParseUint(string(maxUint[:i]), 10, 64)
		bench = append(bench, f32Bench{name, float32(num)})
	}
	for _, c := range bench {
		libs := []struct {
			name string
			test func(*testing.B)
		}{{
			name: "StdLib",
			test: func(b *testing.B) {
				_, _ = json.Marshal(c.float)
				for i := 0; i < b.N; i++ {
					_, _ = json.Marshal(c.float)
				}
			},
		}, {
			name: "Sonic",
			test: func(b *testing.B) {
				_, _ = encoder.Encode(c.float, 0)
				for i := 0; i < b.N; i++ {
					_, _ = encoder.Encode(c.float, 0)
				}
			},
		}}
		for _, lib := range libs {
			name := lib.name + "_" + c.name
			b.Run(name, lib.test)
		}
	}
}
