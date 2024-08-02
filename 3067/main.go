package main

import "fmt"

func main() {
	fmt.Println(countPairsOfConnectableServers([][]int{
		{0,1,1},
		{1,2,5},
		{2,3,13}, 
		{3,4,9},
		{4,5,2},
	}, 1))
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	graph := make([][][]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], []int{v, w})
		graph[v] = append(graph[v], []int{u, w})
	}

	var dfs func(int, int, int) int
	dfs = func(p, root, curr int) int {
		res := 0
		if curr == 0 {
			res++
		}

		for _, e := range graph[p] {
			v, cost := e[0], e[1]
			if v != root {
				res += dfs(v, p, (curr + cost) % signalSpeed)
			}
		}
		return res
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		pre := 0
		for _, e := range graph[i] {
			v, cost := e[0], e[1]
			cnt := dfs(v, i , cost %signalSpeed)
			res[i] += pre * cnt
			pre += cnt
		}
	}
	return res
}