package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOfPairs([]int{5,5,5,5}))
}

func countOfPairs(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 51)
	}
	mod := 1000000007
	for v := 0; v <= nums[0]; v++{
		dp[0][v] = 1
	}
	for i := 1; i < n; i++ {
		for v2 := 0; v2 <= nums[i]; v2++ {
			for v1 := 0; v1 <= v2; v1++ {
				if nums[i-1]-v1 >= nums[i]-v2 && nums[i]-v2 >= 0 {
					dp[i][v2] = (dp[i][v2] + dp[i-1][v1] % mod)
				}
			}
		}
	}
	res := 0
	for _, v := range dp[n-1] {
		res = (res + v) % mod
	}
	return res
}