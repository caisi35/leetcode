package main

import "fmt"

func main() {
	fmt.Println(minRefuelStops(1, 1, [][]int{}))
}

func minRefuelStops(target, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i, s := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= s[0] {
				dp[j+1] = max(dp[j+1], dp[j] + s[1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}