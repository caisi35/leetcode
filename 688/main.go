package main

import (
	"fmt"
)

func main() {
	fmt.Println(knightProbability(1,0,0,0))
}

func knightProbability(n, k, row, column int) float64 {
	dirs := []struct{i, j int}{
		{-2,-1},
		{-2,1},
		{2,-1},
		{2,1},
		{-1,-2},
		{1,-2},
		{-1,2},
		{1,2},
	}
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dirs {
						if x, y := i+d.i, j+d.j; 0 <= x && x < n && 0 <= y && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}