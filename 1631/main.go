package main

import (
	"fmt"
	"sort"
)

func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)	// 0 0 0 0 0 1 2 3

	s2 := make([]int,0)
	s2 = append(s2,1,2,3,4)
	fmt.Println(s2) // 1 2 3 4
	fmt.Println(minimumEffortPath([][]int{
		{1,2,1,1,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,1,1,2,1},
	}))
}

type pair struct {
	x, y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range dirs {
				x, y := p.x + d.x, p.y + d.y
				if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && abs(heights[x][y] - heights[p.x][p.y]) <= maxHeightDiff {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

func abs(i int) int{
	if i < 0 {
		return -i
	}
	return i
}