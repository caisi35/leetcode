package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubarrays([]int{1,2,1,4,1})) // 1
	fmt.Println(countSubarrays([]int{0,0,0,0})) // 2
	fmt.Println(countSubarrays([]int{-1,-4,-1,4})) // 1
}

func countSubarrays(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 2; i < n; i++ {
		if (nums[i] + nums[i-2])*2 == nums[i-1] {
			ans++
		}
	}
	return ans
}