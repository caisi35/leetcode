package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxConsecutive(2, 9, []int{4,6}))		// 3
	fmt.Println(maxConsecutive(6, 8, []int{7,6,8}))		// 0
	fmt.Println(maxConsecutive(3, 15, []int{7,9,13}))	// 4
	fmt.Println(maxConsecutive(28, 50, []int{35, 48}))	// 12
}

func maxConsecutive(bottom int, top int, special []int) int {
	ans := 0
	if len(special) == (top - bottom + 1) {
		return ans
	}

	sort.Ints(special)

	for i := 1; i < len(special); i++ {
		ans = max(ans, special[i] - special[i-1] - 1)
	}
	ans = max(ans, special[0] - bottom, top - special[len(special)-1])
	return ans
}