package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Println(maxProfitAssignment([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}))

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	dp := [][]int{}
	ans := 0
	for i, d := range difficulty {
		dp = append(dp, []int{d, profit[i]})
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i][0] < dp[j][0]
	})
	sort.Ints(worker)
	// fmt.Println(dp)
	// fmt.Println(worker)
	for _, w := range worker {
		m := 0
		for _, d := range dp {
			if w >= d[0] {
				m = max(m, d[1])
			}
		}
		ans += m
	}
	return ans
}
