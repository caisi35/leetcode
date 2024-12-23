package main

import (
	"fmt"
)

func main() {
	var a = []int{1,2,3,4,5}
	var r [5]int
	for i, v := range a {
		fmt.Println(i, v, a)
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println(findRedundantDirectedConnection([][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}))

}

func findRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	n := len(edges)
	uf := newUnionFind(n + 1)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to {
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) {
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}
	if conflictEdge == nil {
		return cycleEdge
	}
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
