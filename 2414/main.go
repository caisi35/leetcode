package main

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacabc"))	// 3
}

func longestContinuousSubstring(s string) int {
	cur := 1
	ans := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] + byte(1) {
			cur++
		} else {
			cur = 1
		}
		ans = max(cur, ans)
	}
	return ans
}