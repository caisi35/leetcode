package main

import "fmt"

type Integer int

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
	fmt.Println(firstDayBeenInAllRooms([]int{0, 1, 2, 0}))
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int(1e9 + 7)
	dp := make([]int, len(nextVisit))
	dp[0] = 2
	for i := 1; i < len(nextVisit); i++ {
		to := nextVisit[i]
		dp[i] = dp[i-1] + 2
		if to != 0 {
			dp[i] = (dp[i] - dp[to-1] + mod) % mod
		}
		dp[i] = (dp[i] + dp[i-1]) % mod
	}
	return dp[len(nextVisit)-2]
}
