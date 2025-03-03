package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCut("abc"))
}

func minCut(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
		for j := range isPalindrome[i] {
			isPalindrome[i][j] = true
		}
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				isPalindrome[i][j] = isPalindrome[i+1][j-1]
			} else {
				isPalindrome[i][j] = false
			}
		}
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if isPalindrome[j][i-1] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n] - 1
}
