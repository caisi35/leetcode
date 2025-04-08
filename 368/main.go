package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestDivisibleSubset([]int{1,2,3,4}))
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i] % v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	if maxSize == 1 {
		return []int{nums[0]}
	}
	res := []int{}
	for i := n-1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal % nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}