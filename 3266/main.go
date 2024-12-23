package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getFinalState([]int{2,1,3,5,6}, 5, 2))
}

func quickMul(x, y, m int64) int64 {
	res := int64(1)
	for y > 0 {
		if ( y & 1) == 1 {
			res = (res * x) % m
		}
		y >>= 1
		x = (x * x) % m
	}
	return res
}

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n, m := len(nums), int64(1e9+7)
	mx := 0
	var v minHeap
	for i, num := range nums {
		mx = max(mx, num)
		v = append(v, pair{int64(num), i})
	}
	heap.Init(&v)
	for ; v[0].first < int64(mx) && k > 0; k-- {
		x := heap.Pop(&v).(pair)
		x.first *= int64(multiplier)
		heap.Push(&v, x)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].first < v[j].first || v[i].first == v[j].first && v[i].second < v[j].second
	})
	for i := 0; i < n; i++ {
		t := k / n
		if i < k % n {
			t++
		}
		nums[v[i].second] = int((v[i].first % m) * quickMul(int64(multiplier), int64(t), m) % m)
	}
	return nums
}

type pair struct {
	first int64
	second int
}

type minHeap []pair

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].first < h[j].first || h[i].first == h[j].first && h[i].second < h[j].second
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[0: n-1]
	return res
}