package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	win := make(map[int]int)
	los := make(map[int]int)
	for _, matche := range matches {
		w, l := matche[0], matche[1]
		win[w]++
		los[l]++
	}
	
	ans0 := []int{}
	for w  := range win {
		if _, ok := los[w]; !ok {
			ans0 = append(ans0, w)
		}
	}

	ans1 := []int{}
	for l, c := range los {
		if c == 1 {
			ans1 = append(ans1, l)
		}
	}
	sort.Ints(ans0)
	sort.Ints(ans1)

	ans := [][]int{}
	ans = append(ans, ans0)
	ans = append(ans, ans1)
	return ans
}

func main() {
	fmt.Println(findWinners([][]int{
		{2, 3}, {1, 7}, {5, 4}, {6, 4},
	}))
	fmt.Println(findWinners([][]int{
		{1,3},{2,3},{3,6},{5,6},{5,7},{4,5},{4,8},{4,9},{10,4},{10,9},
	}))
}
