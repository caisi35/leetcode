package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countWays([][]int{
		{1, 3}, {10, 20}, {2, 5}, {4, 8},
	}))
	fmt.Println(countWays([][]int{
		{6, 10}, {5, 15},
	}))
	fmt.Println(countWays([][]int{
		{0, 0}, {8, 9}, {12, 13}, {1, 3},
	}))
}

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	
	res := 1
	n := len(ranges)
	for i := 0; i < n; {
		r := ranges[i][1]
		j := i+1
		for j < n && ranges[j][0] <= r {
			r = max(r, ranges[j][1])
			j++
		}
		res = (res*2)
		i = j
	}
	return res
}
