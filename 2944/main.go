package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCoins([]int{3,1,2}))
}

func minimumCoins(prices []int) int {
	memo := make(map[int]int)

	var dp func(index int) int 
	dp = func(index int) int {
		if 2 * index + 2 >= len(prices) {
			return prices[index]
		}
		if val, ok := memo[index]; ok {
			return val
		}
		minValue := math.MaxInt32
		for i := index + 1; i <= 2 * index + 2; i++ {
			minValue = min(minValue, dp(i))
		}
		memo[index] = prices[index] + minValue
		return memo[index]
	}
	return dp(0)
}