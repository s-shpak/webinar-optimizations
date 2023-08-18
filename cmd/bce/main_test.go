package main

import (
	"log"
	"testing"
)

func TestSumFirstElementsOfSlice(t *testing.T) {
	s := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	res := SumFirstElementsOfSlice(s, 100_000_000)
	log.Println(res)
}

func BenchmarkSumFirstElementsOfSlice(b *testing.B) {
	s := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumFirstElementsOfSlice(s, 100000)
	}
}

func BenchmarkSumFirstElementsOfSliceBCE(b *testing.B) {
	s := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumFirstElementsOfSliceBCE(s, 100000)
	}
}
