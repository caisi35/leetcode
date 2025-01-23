package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2,4,1,2,7,8}))	// 9
}

func maxCoins(piles []int) int {
	ans := 0 
	sort.Ints(piles)
	left, right := 0, len(piles)-1
	for left <= right {
		ans += piles[right-1]
		left++
		right -= 2
	}
	return ans
}