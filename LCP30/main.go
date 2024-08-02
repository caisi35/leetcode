package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(magicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}))
}

type PriorityQueue struct {
	sort.IntSlice
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.IntSlice[i] < pq.IntSlice[j]
}

func (pq *PriorityQueue) Push(v interface{}) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	arr := pq.IntSlice
	v := arr[len(arr)-1]
	pq.IntSlice = arr[:len(arr)-1]
	return v
}

func magicTower(nums []int) int {
	q := &PriorityQueue{}
	ans, hp, delay := 0, int64(1), int64(0)
	for _, num := range nums {
		if num < 0 {
			heap.Push(q, num)
		}
		hp += int64(num)
		if hp <= 0 {
			ans++
			delay += int64(q.IntSlice[0])
			hp -= int64(heap.Pop(q).(int))
		}
	}
	if hp + delay <= 0 {
		return -1
	}
	return ans
}
