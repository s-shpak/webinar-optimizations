package main

func SumFirstElementsOfSlice(s []int, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += s[i]
	}
	return sum
}
