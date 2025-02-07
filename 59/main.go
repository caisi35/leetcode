package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for num <= n * n {
		for i := left; i <= right; i++ {
			m[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			m[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			m[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			m[i][left] = num
			num++
		}
		left++
	}
	return m
}