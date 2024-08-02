package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheLongestBalancedSubstring("0010"))
}

func findTheLongestBalancedSubstring(s string) int {
	ans := 0
	n := len(s)
	if n < 2 {
		return ans
	}
	z := 0
	o := 0
	for i := 0; i < n; i++ {
		
		if s[i] == '0' {
			if o != 0 {
				t := min(z, o)
				ans = max(t + t, ans)
				z, o = 0, 0
			}
			z++
		} else if s[i] == '1' && z != 0 {
			o++
		}
	}
	if z != 0 || o != 0 {
		t := min(z, o)
		ans = max(t + t, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}