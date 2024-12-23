package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestRangeII([]int{1,3,6}, 3))	// 3
}

func smallestRangeII(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	r := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		a, b := nums[i], nums[i+1]
		r = min(r, max(nums[n-1]-k, a+k) - min(nums[0]+k, b-k))
	}
	return r
}