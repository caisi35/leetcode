package main

import (
	"fmt"
)

func main() {
	fmt.Println(findBall([][]int{{-1}}))	// [-1]
	fmt.Println(findBall([][]int{
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1},
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1}}))		// [0,1,2,3,4,-1]
}

func findBall(grid [][]int) []int {
	ans := make([]int, len(grid[0]))

	for g := range grid[0] {
		colume := g
		for i := range len(grid) {
			c := grid[i][colume]
			colume += c
			if colume < 0 || colume == len(ans) || grid[i][colume] != c {
				colume = -1
				break
			}
		}
		ans[g] = colume

	}

	return ans
}
