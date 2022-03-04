package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		A()
	}
}

func BenchmarkB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		B()
	}
}

func BenchmarkC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C()
	}
}

func BenchmarkB1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		B1()
	}
}
