package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestNumber([]int{-4,-2,1,4,8}))
	fmt.Println(findClosestNumber([]int{2,-1,1}))
	fmt.Println(findClosestNumber([]int{-1,-1}))
}

func findClosestNumber(nums []int) int {
	ans := 100000000
	for i := range nums {
		n := abs(nums[i] - 0)
		if n < abs(ans) {
			ans = nums[i]
		}
		if abs(n) == abs(ans) {
			ans = max(nums[i], ans)
		}
	}
	return ans 
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}