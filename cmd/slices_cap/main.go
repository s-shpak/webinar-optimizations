package main

func CreateSliceNoCap(i int) []int {
	s := make([]int, 0)
	for j := 0; j <= i; j++ {
		s = append(s, j)
	}
	return s
}

func CreateSliceWithCap(i int) []int {
	s := make([]int, 0, i+1)
	for j := 0; j <= i; j++ {
		s = append(s, j)
	}
	return s
}
