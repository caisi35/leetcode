package main

import (
	"fmt"
	"sort"
)
func main() {
	fmt.Println(numberGame([]int{5,4,2,3}))
}

func numberGame(nums []int) []int {
	ans := make([]int, len(nums))
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		ans[i] = nums[i-1]
		ans[i-1] = nums[i]
	}
	return ans
}