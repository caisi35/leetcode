package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(maximumDetonation([][]int{
		{2,1,3},
		{6,1,4},
	}))
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	isConn := func(u, v int) bool {
		dx := bombs[u][0] - bombs[v][0]
		dy := bombs[u][1] - bombs[v][1]
		return int64(bombs[u][2] * bombs[u][2]) >= int64(dx*dx + dy*dy)
	}

	edges := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConn(i, j) {
				edges[i] = append(edges[i], j)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		visited := make([]int, n)
		cnt := 1
		q := list.New()
		q.PushBack(i)
		visited[i] = 1
		for q.Len() > 0 {
			cidx := q.Remove(q.Front()).(int)
			for _, nidx := range edges[cidx] {
				if visited[nidx] != 0 {
					continue
				}
				cnt++
				q.PushBack(nidx)
				visited[nidx] = 1
			}
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}