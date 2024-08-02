package main

import (
	"fmt"
)

func main() {
	// var x string = nil
    // if x == nil {
    //     x = "default"
    // }
    // fmt.Println(x)
	
	fmt.Println(minIncrements(7, []int{1,5,2,2,3,3,1}))
}

func minIncrements(n int, cost []int) int {
	ans := 0
	for i := n - 2; i > 0; i -= 2 {
		ans += abs(cost[i] - cost[i+1])
		cost[i / 2] = cost[i / 2] + max(cost[i], cost[i + 1])
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}