package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumTripletValue([]int{12,6,1,2,7}))
}

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], nums[i-1])
		right[n-1-i] = max(right[n-i], nums[n-i])
	}
	ans := int64(0)
	for j := 1; j < n - 1; j++ {
		ans = max(ans, int64((left[j] - nums[j]) * right[j]))
	}
	return ans
}