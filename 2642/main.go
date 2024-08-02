package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	type MyInt int
	var i int = 1

	// var a MyInt = i

	// var d MyInt = (MyInt)i

	var b MyInt = MyInt(i)

	// var c MyInt = i.(MyInt)
	fmt.Println(b)

	obj := Constructor(4, [][]int{
		{0, 2, 5},
		{0, 1, 2},
		{1, 2, 1},
		{3, 0, 3},
	})
	fmt.Println(obj.ShortestPath(3, 2))
	fmt.Println(obj.ShortestPath(0, 3))
	obj.AddEdge([]int{1, 3, 4})
	fmt.Println(obj.ShortestPath(0, 3))

}

type Graph struct {
	graph [][]Pair
}

func Constructor(n int, edges [][]int) Graph {
	g := make([][]Pair, n)
	for i := 0; i < n; i++ {
		g[i] = []Pair{}
	}
	for _, e := range edges {
		x, y, cost := e[0], e[1], e[2]
		g[x] = append(g[x], Pair{y, cost})
	}
	return Graph{g}
}

func (this *Graph) AddEdge(edge []int) {
	x, y, cost := edge[0], edge[1], edge[2]
	(*this).graph[x] = append((*this).graph[x], Pair{y, cost})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	pq := PriorityQueue{}
	dist := make([]int, len((*this).graph))
	for i := 0; i < len((*this).graph); i++ {
		dist[i] = math.MaxInt32
	}
	dist[node1] = 0
	heap.Push(&pq, Pair{0, node1})
	for len(pq) > 0 {
		p := heap.Pop(&pq).(Pair)
		cost, cur := p.first, p.second
		if cur == node2 {
			return cost
		}
		for _, e := range (*this).graph[cur] {
			next, ncost := e.first, e.second
			if dist[next] > cost+ncost {
				dist[next] = cost + ncost
				heap.Push(&pq, Pair{cost + ncost, next})
			}
		}
	}
	return -1
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
