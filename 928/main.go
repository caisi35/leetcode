package main

import (
	"fmt"
)

func dfs(grap [][]int, initSet , infeSet []int, v int) {
	n := len(grap)
	for u := 0; u < n; u++ {
		if grap[v][u] == 0 || initSet[u] == 1 || infeSet[u] == 1 {
			continue
		}
		infeSet[u] = 1
		dfs(grap, initSet, infeSet, u)
	}
}

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	initSet := make([]int, n)
	for _, v := range initial {
		initSet[v] = 1
	}
	infeBy := make([][]int, n)
	for _, v := range initial {
		infeSet := make([]int, n)
		dfs(graph, initSet, infeSet, v)
		for u := 0; u < n; u++ {
			if infeSet[u] == 1 {
				infeBy[u] = append(infeBy[u], v)
			}
		}
	}
	count := make([]int, n)
	for u := 0; u < n; u++ {
		if len(infeBy[u]) == 1 {
			count[infeBy[u][0]]++
		}
	}
	res := initial[0]
	for _, v := range initial {
		if count[v] > count[res] || count[v] == count[res] && v < res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(minMalwareSpread([][]int{
		{1,1,0},
		{1,1,1},
		{0,1,1},
	}, []int{0, 1}))
}