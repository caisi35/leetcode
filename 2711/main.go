package main

import (
	"fmt"
)

func main() {
	fmt.Println(differenceOfDistinctValues([][]int{{1,2,3},{3,1,5},{3,2,1}}))
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range m {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			leftTop := make(map[int]struct{})
			ki, kj := i-1, j-1
			for ki >= 0 && kj >= 0 {
				leftTop[grid[ki][kj]] = struct{}{}
				ki--
				kj--
			}
			right := make(map[int]struct{})
			ki, kj = i+1, j+1
			for ki < m && kj < n {
				right[grid[ki][kj]] = struct{}{}
				ki++
				kj++
			}
			ans[i][j] = abs(len(right)-len(leftTop))
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}