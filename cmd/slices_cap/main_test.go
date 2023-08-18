package main

import "testing"

func BenchmarkCreateSliceNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateSliceNoCap(i)
	}
}

func BenchmarkCreateSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateSliceWithCap(i)
	}
}

func TestPrintSliceCapacityChanges(t *testing.T) {
	var intSlice []int

	oldCapacity := cap(intSlice)
	for i := 0; i < 2000; i++ {
		intSlice = append(intSlice, i)
		newCapacity := cap(intSlice)
		if oldCapacity != newCapacity {
			t.Logf("old capacity: %d, new capacity: %d", oldCapacity, newCapacity)
			oldCapacity = newCapacity
		}
	}
}
