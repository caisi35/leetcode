package main

import (
	"fmt"
)

func main() {
	fmt.Println(findIntegers(5))
}

func findIntegers(n int) int {
	ans := 0
	dp := [31]int{1,1}
	// fmt.Println(dp)
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	for i, pre := 29,0; i >= 0; i-- {
		val := 1 << i
		if n & val > 0 {
			ans += dp[i+1]
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		if i == 0 {
			ans++
		}
	}
	return ans
}