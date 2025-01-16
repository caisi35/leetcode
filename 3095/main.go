package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSubarrayLength([]int{1,2,3}, 2))		// 1
	fmt.Println(minimumSubarrayLength([]int{8,2,1}, 10))	// 3
	fmt.Println(minimumSubarrayLength([]int{32,2,24,1}, 35))	// 3
}

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	res := n+1
	for i := 0; i < n; i++ {
		v := 0
		for j := i; j < n; j++ {
			v |= nums[j]
			if v >= k {
				res = min(res, j-i+1)
				break
			}
		}
	}
	if res == n + 1{
		return -1
	}
	return res
}