package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumCost("010101"))
}

func minimumCost(s string) int64 {
	ans := int64(0)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, len(s)-i))
		}
	}
	return ans
}