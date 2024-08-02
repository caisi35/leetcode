package main

import "fmt"

func main() {
	fmt.Println(c([]int{1, 2, 3}, 4))
}

func c(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}