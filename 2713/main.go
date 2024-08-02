package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxIncreasingCells([][]int{
		{3,1},
		{3,4},
	}))
}

func maxIncreasingCells(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	mp := make(map[int][][2]int)
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[mat[i][j]] = append(mp[mat[i][j]], [2]int{i, j})
		}
	}
	keys := []int{}
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		pos := mp[key]
		res := []int{}
		for _, arr := range pos {
			res = append(res, max(row[arr[0]], col[arr[1]]) + 1)
		}
		for i := 0; i < len(pos); i++ {
			arr, d := pos[i], res[i]
			row[arr[0]] = max(row[arr[0]], d)
			col[arr[1]] = max(col[arr[1]], d)
		}
	}

	return maxSlice(row)
}

func maxSlice(slice []int) int {
	maxVal := slice[0]
	for _, val := range slice {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}