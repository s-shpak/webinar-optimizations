package main

import "testing"

func BenchmarkFormatIntStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatIntStrconv(i)
	}
}

func BenchmarkFormatIntSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormatIntSprintf(i)
	}
}
