package main

import (
	"fmt"
	"math"
	"math/bits"
)


func main() {
	list := new([]int)
	// list = append(list, 1)
	fmt.Println(list)
	fmt.Println(maxStudents([][]byte{
		{'#','.','#','#','.','#'},
		{'.','#','#','#','#','.'},
		{'#','.','#','#','.','#'},
	}))
}

func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	memo := make(map[int]int)

	isSingleRowCompliant := func(status, row int) bool {
		n := len(seats[0])
		for j := 0; j < n; j++ {
			if (status >> j) & 1 == 1 {
				if seats[row][j] == '#' {
					return false
				}
				if j > 0 && (status >> (j - 1)) & 1 == 1 {
					return false
				}
			} 
		}
		return true
	}
	isCrossRowsCompliant := func(status, upperRowStatus int) bool {
		n := len(seats[0])
		for j := 0; j < n; j++ {
			if (status >> j) & 1 == 1 {
				if j < n - 1 && (upperRowStatus >> (j + 1)) & 1 == 1 {
					return false
				}
				if j > 0 && (upperRowStatus >> (j - 1)) & 1 == 1 {
					return false
				}
			}
		}
		return true
	}
	var dp func(int, int) int
	dp = func(row, status int) int {
		n := len(seats[0])
		key := (row << n) + status
		if _, ok := memo[key]; !ok {
			if !isSingleRowCompliant(status, row) {
				memo[key] = math.MinInt32
				return math.MinInt32
			}
			students := bits.OnesCount(uint(status))
			if row == 0 {
				memo[key] = students
				return students
			}
			mx := 0
			for upperRowStatus := 0; upperRowStatus < 1 << n; upperRowStatus++ {
				if isCrossRowsCompliant(status, upperRowStatus) {
					mx = max(mx, dp(row - 1, upperRowStatus))
				}
			}
			memo[key] = students + mx
		}
		return memo[key]
	}
	mx := 0
	for i := 0; i < (1 << n); i++ {
		mx = max(mx, dp(m - 1, i))
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
