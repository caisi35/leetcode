package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countWays([]int{1,1})) // 2
}
func countWays(nums []int) int {
	n := len(nums)
	res := 0 
	sort.Ints(nums)
	for k := 0; k <= n; k++ {
		if k > 0 && nums[k-1] >= k {
			continue
		}
		if k < n && nums[k] <= k {
			continue
		}
		res++
	}
	return res
}