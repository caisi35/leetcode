package main

import "fmt"

func main() {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
	fmt.Println(minMalwareSpread([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}, []int{0, 1}))
}

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	ids := make(map[int]int)
	idToSize := make(map[int]int)
	id := 0
	for i := range ids {
		if ids[i] == 0 {
			id++
			ids[i] = id
			size := 1
			q := []int{i}
			for len(q) > 0 {
				u := q[0]
				q = q[1:]
				for v := range graph[u] {
					if ids[v] == 0 && graph[u][v] == 1 {
						size++
						q = append(q, v)
						ids[v] = id
					}
				}
			}
			idToSize[id] = size
		}
	}
	idToInit := make(map[int]int)
	for _, u := range initial {
		idToInit[ids[u]]++
	}
	ans := n + 1
	ansRemoved := 0
	for _, u := range initial {
		removed := 0
		if idToInit[ids[u]] == 1 {
			removed = idToSize[ids[u]]
		}
		if removed > ansRemoved || (removed == ansRemoved && u < ans) {
			ans, ansRemoved = u, removed
		}
	}
	return ans
}
