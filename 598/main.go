package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2,2},{3,3}}))	// 4
}

func maxCount(m, n int, ops [][]int) int {
	a, b := m, n
	for _, op := range ops {
		a = min(a, op[0])
		b = min(b, op[1])
	}
	return a * b
}