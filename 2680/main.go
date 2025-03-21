package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumOr([]int{12,9}, 1))
}

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int64, n+1)
	for i := n-1; i >= 0; i-- {
		suf[i] = suf[i+1] | int64(nums[i])
	}
	var res, pre int64
	for i := 0; i < n; i++ {
		res = max(res, pre | (int64(nums[i]) << k) | suf[i+1])
		pre |= int64(nums[i])
	}
	return res
}