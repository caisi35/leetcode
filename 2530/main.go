package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type PriorityQueue [] int

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	q := (*PriorityQueue)(&nums)
	heap.Init(q)
	var ans int64

	for i := 1; i <= k; i++ {
		x := heap.Pop(q).(int)
		ans += int64(x)
		heap.Push(q, (x + 2) / 3)
	}
	return ans
}

func main() {
	GuessingGame()
	fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5))
}

func GuessingGame() {
    var s string
    fmt.Printf("Pick an integer from 0 to 100.\n")
    answer := sort.Search(100, func(i int) bool {
        fmt.Printf("Is your number <= %d? ", i)
        fmt.Scanf("%s", &s)
        return s != "" && s[0] == 'y'
    })
    fmt.Printf("Your number is %d.\n", answer)
}