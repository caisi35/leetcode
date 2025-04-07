package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartition([]int{1,5,11,5}))
	fmt.Println(canPartition([]int{2,2,1,1}))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	m := sum / 2
	dp := make([]bool, m + 1)
	dp[0] = true
	for _, num := range nums {
		for j := m; j >= num; j-- {
			if dp[j-num] {
				dp[j] = true
			}
		}
	}
	return dp[m]
}