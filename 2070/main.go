package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{10,1000}}, []int{5}))
}

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	n := len(items)
	for i := 1; i < n; i++ {
		if items[i][1] < items[i-1][1] {
			items[i][1] = items[i-1][1]
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = query(items, q)
	}
	return res
}

func query(items [][]int, q int) int {
	l, r := 0, len(items)
	for l < r {
		mid := l + (r - l) / 2
		if items[mid][0] > q {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		return 0
	} else {
		return items[l-1][1]
	}
}