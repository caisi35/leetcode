package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(palindromePartition("abcs", 2))
}

func palindromePartition(s string, k int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for l := 0; l < i; l++ {
				cost := calCost(s, l, i-1)
				if dp[l][j-1] != math.MaxInt32 {
					dp[i][j] = min(dp[i][j], dp[l][j-1] + cost)
				}
			}
		}
	}
	if dp[n][k] == math.MaxInt32 {
		return -1
	}
	return dp[n][k]
}

func calCost(s string, l, r int) int {
	cost := 0 
	for l < r {
		if s[l] != s[r] {
			cost++
		}
		r--
		l++
	}
	return cost
}