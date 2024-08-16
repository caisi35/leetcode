package main

import (
	"fmt"
	"math"
	"regexp"
)

func main() {
	fmt.Println(maxScore([][]int{
		{9,5,7,3},
		{8,9,6,1},
		{6,7,14,3},
		{2,5,3,1},
	}))
	re, err := regexp.Compile(fmt.Sprintf(`^#{1,2} %s / (\d{4}-\d{2}-\d{2})`, regexp.QuoteMeta("v1.16.03")))
	fmt.Println(re, err)
	re = regexp.MustCompile(`^\* \[([^\]]+)\]`)
	fmt.Println(re)
}

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	premin := make([][]int, 2)
	for i := range premin {
		premin[i] = make([]int, n)
		for j := range premin[i] {
			premin[i][j] = math.MaxInt
		}
	}
	ans := math.MinInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pre := math.MaxInt
			if i > 0 {
				pre = min(pre, premin[(i-1) & 1][j])
			}
			if j > 0 {
				pre = min(pre, premin[i & 1][j-1])
			}
			if i + j > 0 {
				ans = max(ans, grid[i][j] - pre)
			}
			premin[i&1][j] = min(pre, grid[i][j])
		}
	}
	return ans
}

func maxScore3(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	premin := make([][]int, m)
	for i := range premin {
		premin[i] = make([]int, n)
		for j := range premin[i] {
			premin[i][j] = math.MaxInt32
		}
	}
	ans := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pre := math.MaxInt32
			if i > 0 {
				pre = min(pre, premin[i-1][j])
			}
			if j > 0 {
				pre = min(pre, premin[i][j-1])
			}
			if i + j > 0 {
				ans = max(ans, grid[i][j] - pre)
			}
			premin[i][j] = min(pre, grid[i][j])
		}
	}
	return ans
}

func maxScore2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	precol := make([]int, n)
	for i := range precol {
		precol[i] = math.MinInt32
	}
	ans := math.MinInt32
	for i := 0; i < m; i++ {
		prerow := math.MinInt32
		for j := 0; j < n; j++ {
			f := math.MinInt32
			if i > 0 {
				f = max(f, grid[i][j]+precol[j])
			}
			if j > 0 {
				f = max(f, grid[i][j] + prerow)
			}
			ans = max(ans, f)
			precol[j] = max(precol[j], max(f, 0) - grid[i][j])
			prerow = max(prerow, max(f, 0) - grid[i][j])
		}
	}
	return ans
}

func maxScore1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	prerow := make([][]int, m)
	precol := make([][]int, m)
	f := make([][]int, m)
	for i := range prerow {
		prerow[i] = make([]int, n)
		precol[i] = make([]int, n)
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	ans := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				f[i][j] = max(f[i][j], grid[i][j] + precol[i-1][j])
			}
			if j > 0 {
				f[i][j] = max(f[i][j], grid[i][j] + prerow[i][j-1])
			}
			ans = max(ans,f[i][j])
			prerow[i][j] = max(f[i][j], 0) - grid[i][j]
			precol[i][j] = prerow[i][j]
			if i > 0 {
				precol[i][j] = max(precol[i][j], precol[i-1][j])
			}
			if j > 0 {
				prerow[i][j] = max(prerow[i][j], prerow[i][j-1])
			}
		}
	}
	return ans
}