package main

import (
	"container/heap"
	"fmt"
)

func main() {
	for range 10 {
		foo()
	}
	fmt.Println(minimumVisitedCells([][]int{
		{3, 4, 2, 1},
		{4, 2, 3, 1},
		{2, 1, 0, 0},
		{2, 4, 0, 0},
	}))
}

func foo() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 1
	row := make([]PriorityQueue, m)
	col := make([]PriorityQueue, n)

	update := func(x *int, y int) {
		if *x == -1 || y < *x {
			*x = y
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for len(row[i]) > 0 && row[i][0].second+grid[i][row[i][0].second] < j {
				heap.Pop(&row[i])
			}
			if len(row[i]) > 0 {
				update(&dist[i][j], dist[i][row[i][0].second]+1)
			}
			for len(col[j]) > 0 && col[j][0].second+grid[col[j][0].second][j] < i {
				heap.Pop(&col[j])
			}
			if len(col[j]) > 0 {
				update(&dist[i][j], dist[col[j][0].second][j]+1)
			}
			if dist[i][j] != -1 {
				heap.Push(&row[i], Pair{dist[i][j], j})
				heap.Push(&col[j], Pair{dist[i][j], j})
			}
		}
	}
	return dist[m-1][n-1]
}

type Pair struct {
	first  int
	second int
}

type PriorityQueue []Pair

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}
