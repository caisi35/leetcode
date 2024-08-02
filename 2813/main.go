package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaximumElegance([][]int{
		{3,2},{5,1},{10,1},
	}, 2))
}

func findMaximumElegance(items [][]int, k int) int64 {
	sort.Slice(items, func(a, b int) bool {
		return items[a][0] > items[b][0]
	})
	ca := map[int]bool{}
	var res, profix int64
	var st []int
	for i, item := range items {
		if i < k {
			profix += int64(item[0])
			if ca[item[1]] {
				st = append(st, item[0])
			} else {
				ca[item[1]] = true
			}
		} else if !ca[item[1]] && len(st) > 0 {
			profix += int64(item[0] - st[len(st)-1])
			st = st[:len(st)-1]
			ca[item[1]] = true
		}
		res = max(res, profix + int64(len(ca) * len(ca)))
	}
	return res
}