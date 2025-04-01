package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostPoints([][]int{{3,2},{4,3},{4,4},{2,5}})) // 5
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i := n-1; i >= 0; i-- {
		next := min(n, i+int(questions[i][1])+1)
		take := int64(questions[i][0]) + dp[next]
		skip := dp[i+1]
		dp[i] = max(take, skip)
	}
	return dp[0]
}