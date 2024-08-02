package main

import (
	"fmt"
)

func main() {
	var a int = 3
	// 以下有额外内存分配吗？
	var i interface{} = a
	fmt.Printf("%v, %v\n", &a, &i)
	fmt.Println(findChampion([][]int{
		{0, 0, 1},
		{1, 0, 1},
		{0, 0, 0},
	}))
}

func findChampion(grid [][]int) int {
	n := len(grid)
	mx := 0
	ans := 0
	for i := 0; i < n; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				c++
			}
		}
		if mx < c {
			mx = c
			ans = i
		}
	}
	return ans
}
