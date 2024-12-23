package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(12, 15, 2))
	fmt.Println(getKth(7, 11, 4))
	fmt.Println(getKth(1, 1000, 777))
}

func getKth(lo int, hi int, k int) int {
	ks := [][]int{}

	s := func(x int) int {
		ans := 0
		for x != 1 {
			if x % 2 == 0 {
				x /= 2
			} else {
				x = 3 * x + 1
			}
			ans++
		}
		return ans
	}
	for i := lo; i <= hi; i++ {
		ks = append(ks, []int{s(i), i})
	}
	sort.Slice(ks, func(i, j int) bool {
		if ks[i][0] == ks[j][0] {
			return ks[i][1] < ks[j][1]
		}
		return ks[i][0] < ks[j][0]
	})
	// fmt.Println(ks)
	return ks[k-1][1]
}