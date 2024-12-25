package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCost(3,2,[]int{1,3}, []int{5}))
}

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	cache := make([]int, m*m*n*n)
	for i := range cache {
		cache[i] = -1
	}

	index := func(row1, col1, row2, col2 int) int {
		return (row1*n + col1) * m * n + row2 * n + col2
	}

	var dp func(row1, col1, row2, col2 int) int
	dp = func(row1, col1, row2, col2 int) int {
		if row1 == row2 && col1 == col2 {
			return 0
		}
		ind := index(row1, col1, row2, col2)
		if cache[ind] >= 0 {
			return cache[ind]
		}
		cache[ind] = math.MaxInt32
		for i := row1; i < row2; i++ {
			cache[ind] = min(cache[ind], dp(row1, col1, i, col2) + dp(i+1, col1, row2, col2) + horizontalCut[i])
		}
		for i := col1; i < col2; i++ {
			cache[ind] = min(cache[ind], dp(row1, col1, row2, i) + dp(row1, i+1, row2, col2) + verticalCut[i])
		}
		return cache[ind]
	}
	return dp(0, 0, m-1, n-1)
}