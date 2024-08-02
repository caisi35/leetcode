package main

import (
	"fmt"
	"sort"
)

func main() {
	a := 5
	b := 8.1
	fmt.Println(float64(a) + b)
	fmt.Println(minimumRemoval([]int{6, 1, 4, 5}))
}

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	total := int64(0)
	for _, bean := range beans {
		total += int64(bean)
	}
	res := total
	for i := 0; i < n; i++ {
		res = min(res, total-int64(beans[i])*int64(n-i))
	}
	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
