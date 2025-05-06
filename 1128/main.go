package main

import (
	"fmt"
)

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))         // 1
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}})) // 3
}

func numEquivDominoPairs(dominoes [][]int) int {
	ans := 0
	cnt := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0] * 10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}
