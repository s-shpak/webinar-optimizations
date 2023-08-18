package main

import (
	"fmt"
	"strconv"
)

func FormatIntStrconv(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatIntSprintf(i int) string {
	return fmt.Sprintf("%d", i)
}
