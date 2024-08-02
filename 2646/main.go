package main

import (
	"fmt"
)

func main() {
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}
	price := []int{2, 2, 10, 6}
	trips := [][]int{{0, 3}, {2, 1}, {2, 3}}
	fmt.Println(minimumTotalPrice(n, edges, price, trips))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	next := make([][]int, n)
	for _, edge := range edges {
		next[edge[0]] = append(next[edge[0]], edge[1])
		next[edge[1]] = append(next[edge[1]], edge[0])
	}

	count := make([]int, n)
	var dfs func(int, int, int) bool
	dfs = func(node, parent, end int) bool {
		if node == end {
			count[node]++
			return true
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			if dfs(child, node, end) {
				count[node]++
				return true
			}
		}
		return false
	}
	for _, trip := range trips {
		dfs(trip[0], -1, trip[1])
	}
	var dp func(int, int) []int
	dp = func(node, parent int) []int {
		res := []int{
			price[node] * count[node], price[node] * count[node] / 2,
		}
		for _, child := range next[node] {
			if child == parent {
				continue
			}
			v := dp(child, node)
			res[0], res[1] = res[0] + min(v[0], v[1]), res[1] + v[0]
		}
		return res
	}
	res := dp(0, -1)
	return min(res[0], res[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}