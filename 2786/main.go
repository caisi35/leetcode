package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxScore([]int{2,3,6,1,9,2}, 5))
}

func maxScore(nums []int, x int) int64 {
	res := int64(nums[0])
	dp := [2]int64{math.MinInt32, math.MinInt32}
	dp[nums[0] % 2] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		par := nums[i] % 2
		cur := max(dp[par] + int64(nums[i]), dp[1 - par] - int64(x) + int64(nums[i]))
		res = max(res, cur)
		dp[par] = max(dp[par], cur)
	}
	return res
}