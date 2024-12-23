package main

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1,4,6,7,8,20}, []int{2,7,15}))
}

func mincostTickets(days []int, costs []int) int {
	memo := [366]int{}
	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}
	var dfs func(day int) int
	dfs = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min(min(dfs(day+1)+costs[0], dfs(day+7)+costs[1]), dfs(day+30)+costs[2])
		} else {
			memo[day] = dfs(day+1)
		}
		return memo[day]
	}
	return dfs(1)
}