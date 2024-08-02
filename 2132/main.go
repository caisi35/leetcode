package main

import "fmt"

func main() {
	m := make(map[string]int)

	fmt.Println(m["qcrao"])
	fmt.Println(possibleToStamp([][]int{
		{1,0,0,0},
		{1,0,0,0},
		{1,0,0,0},
		{1,0,0,0},
	}, 2, 2))
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	psum := make([][]int, m + 2)
	diff := make([][]int, m+2)
	for i:=0; i <= m + 1; i++ {
		psum[i] = make([]int, n + 2)
		diff[i]= make([]int, n+ 2)
	}
	for i := 1; i <= m ; i++ {
		for j := 1; j <= n; j++ {
			psum[i][j] = psum[i-1][j] + psum[i][j-1] - psum[i-1][j-1] + grid[i-1][j-1]
		}
	}
	for i := 1 ; i + stampHeight - 1 <=m;i++ {
		for j := 1; j + stampWidth - 1 <= n; j++ {
			x,  y := i + stampHeight - 1, i + stampWidth - 1
			if psum[x][y] - psum[x][j-1] - psum[i-1][y] + psum[i-1][j-1] == 0 {
				diff[i][j]++
				diff[i][y+1]--
				diff[x+1][j]--
				diff[x+1][y+1]++
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
			if diff[i][j] == 0 && grid[i-1][j-1] == 0 {
				return false
			}
		}
	}
	return true
}