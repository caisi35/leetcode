package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalPathQuality([]int{0, 32, 10, 43}, [][]int{
		{0, 1, 10},
		{1, 2, 15},
		{0, 3, 10},
	}, 49))
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	g := make([][][2]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], [2]int{edge[1], edge[2]})
		g[edge[1]] = append(g[edge[1]], [2]int{edge[0], edge[2]})
	}

	visited := make([]bool, n)
	visited[0] = true
	ans := 0

	var dfs func(int, int, int)
	dfs = func(u, time, value int) {
		if u == 0 {
			ans = max(ans, value)
		}
		for _, e := range g[u] {
			v, dist := e[0], e[1]
			if time + dist <= maxTime {
				if !visited[v] {
					visited[v] = true
					dfs(v, time + dist, value + values[v])
					visited[v] = false
				} else {
					dfs(v, time + dist, value)
				}
			}
		}
	}
	dfs(0, 0, values[0])
	return ans
}