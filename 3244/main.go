package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2,4},{0,2},{0,4}}))
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	roads := make([]int, n)
	for i := 0; i < n; i++ {
		roads[i] = i + 1
	}
	var res []int
	dist := n-1
	for _, query := range queries {
		k := roads[query[0]]
		roads[query[0]] = query[1]
		for k != -1 && k < query[1] {
			k, roads[k] = roads[k], -1
			dist--
		}
		res = append(res, dist)
	}
	return res
}