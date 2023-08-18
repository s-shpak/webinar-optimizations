package main

func SumFirstElementsOfSliceBCE(s []int, n int) int {
	_ = s[n-1]
	var sum int
	for i := 0; i < n; i++ {
		sum += s[i]
	}
	return sum
}
