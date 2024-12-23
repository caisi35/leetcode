package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRedundantConnection([][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}))
}

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(fron, to int) bool {
		x, y := find(fron), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}