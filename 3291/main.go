package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc"))
}

func minValidStrings(words []string, target string) int {
	prefixFunc := func(word, target string) []int {
		s := word + "#" + target
		n := len(s)
		pi := make([]int, n)
		for i := 1; i < n; i++ {
			j := pi[i-1]
			for j > 0 && s[i] != s[j] {
				j = pi[j-1]
			}
			if s[i] == s[j] {
				j++
			}
			pi[i] = j
		}
		return pi
	}
	n := len(target)
	back := make([]int, n)
	for _, word := range words {
		pi := prefixFunc(word, target)
		m := len(word)
		for i := 0; i < n; i++ {
			back[i] = int(math.Max(float64(back[i]), float64(pi[m+1+i])))
		}
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = int(1e9)
	}
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i+1-back[i]] + 1
		if dp[i+1] > n {
			return -1
		}
	}
	return dp[n]
}