package main

import (
	"fmt"
)

func main() {
	fmt.Println(countGoodNodes([][]int{
		{0,1},
		{0,2},
		{1,3},
		{1,4},
		{2,5},
		{2,6},
	}))
}

func countGoodNodes(edges [][]int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	res := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		valid := true
		treeSize := 0
		subTreeSize := 0
		for _, child := range g[node] {
			if child != parent {
				size := dfs(child, node)
				if subTreeSize == 0 {
					subTreeSize = size
				} else if size != subTreeSize {
					valid = false
				}
				treeSize += size
			}
		}
		if valid {
			res++
		}
		return treeSize + 1
	}
	dfs(0, -1)
	return res
}