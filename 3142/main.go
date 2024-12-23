package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(satisfiesConditions([][]int{
		{1,0,2},
		{1,0,2},
	}))		// true
	fmt.Println(satisfiesConditions([][]int{
		{1,1,1},
		{0,0,0},
	}))		// false
	fmt.Println(satisfiesConditions([][]int{
		{1},
		{2},
		{3},
	}))		// false
	foo()
}

func foo() {
	closed := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(closed)
	}()
	select {
	case <- closed:
	case <- time.After(10 + 5 * time.Second):
		print("timeout")
	}
}

func satisfiesConditions(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ((i < n-1) && !(grid[i][j] == grid[i+1][j])) || ((j < m-1) && !(grid[i][j] != grid[i][j+1])) {
				return false
			}
		}
	}
	return true
}