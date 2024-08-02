package main

import (
	"fmt"
	"sort"
)

func main () {
	fmt.Println(minimumDistance([][]int{
		{3, 10},
		{5, 15},
		{10, 2},
		{4, 4},
	}))
}

func minimumDistance(points [][]int) int {
	n := len(points)
	sx := make([][]int, n)
	sy := make([][]int, n)
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		sx[i] = []int{x-y, i}
		sy[i] = []int{x+y, i}
	}
	sort.Slice(sx, func(i, j int) bool {
		return sx[i][0] < sx[j][0]
	})
	sort.Slice(sy, func(i, j int) bool {
		return sy[i][0] < sy[j][0]
	})
	maxVal1 := sx[n-1][0] - sx[0][0]
	maxVal2 := sy[n-1][0] - sy[0][0]
	res := int(^uint(0)>>1)
	if maxVal1 >= maxVal2 {
		i, j := sx[0][1], sx[n-1][1]
		res = min(res, max(remove(sx, i), remove(sy, i)))
		res = min(res, max(remove(sx, j), remove(sy, j)))
	} else {
		i, j := sy[0][1], sy[n-1][1]
		res = min(res, max(remove(sx, i), remove(sy, i)))
		res = min(res, max(remove(sx, j), remove(sy, j)))
	}
	return res
}

func remove(arr [][]int, i int) int {
	n := len(arr)
	if arr[0][1] == i {
		return arr[n-1][0] - arr[1][0]
	} else if arr[n-1][1] == i {
		return arr[n-2][0] - arr[0][0]
	} else {
		return arr[n-1][0]- arr[0][0]
	}
}