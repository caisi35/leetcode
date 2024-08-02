package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{
		{2,1},
		{1, 0},
		{1, 4},
		{1, 8},
		{3, 5},
		{4, 6},
	}, 1))
}

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	bound := -1
	ans := 0 
	for _, point := range points {
		if point[0] > bound {
			bound = point[0] + w
			ans++
		}
	}
	return ans
}