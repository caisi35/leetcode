package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips([][]int{
		{1,0,0},
		{0,1,0},
		{0,0,1},
	}))
}

func minFlips(grid [][]int) int {
	m, n, ans := len(grid), len(grid[0]), 0
	for i := 0; i < m / 2; i++ {
		for j := 0; j < n / 2; j++ {
			cnt1 := grid[i][j] + grid[i][n-1-j] + grid[m-1-i][j] + grid[m-1-i][n-1-j]
			ans += min(cnt1, 4-cnt1)
		}
	}

	diff, cnt1 := 0, 0
	if m % 2 == 1 {
		for j := 0; j < n / 2; j++ {
			if grid[m / 2][j] ^ grid[m / 2][n - 1 - j] != 0 {
				diff++
			} else {
				cnt1 += grid[m / 2][j] * 2
			}
		}
	}
	if n % 2 == 1 {
		for i := 0; i < m / 2; i++ {
			if grid[i][n/2] ^ grid[m-1-i][n/2] != 0 {
				diff++
			} else {
				cnt1 += grid[i][n/2] * 2
			}
		}
	}
	if m % 2 == 1 && n % 2 == 1 {
		ans += grid[m/2][n/2]
	}
	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return ans
}