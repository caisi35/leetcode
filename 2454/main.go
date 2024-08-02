package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(secondGreaterElement([]int{2, 4, 0, 9, 6}))
	fmt.Println(s2([]int{2, 4, 0, 9, 6}))

}

func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	n := len(nums)
	for i := range ans {
		ans[i] = -1
	}
	if n <= 3 {
		return ans
	}

	for i := 0; i < n - 1; i++ {
		f := false
		for j, v := range nums[i+1:] {
			k := j + i + 1
			if k > i && nums[k] > nums[i] {
				if f {
					ans[i] = v
					break
				}
				f = true
				continue
			}
		}
	}

	return ans
}

type Pair struct {
	first int
	second int
}

type hp []Pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.first < b.first || a.first == b.first && a.second < b.second
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(Pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a) - 1]
	*h = a[:len(a)-1]
	return v
}

func s2(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}

	st := []int{}
	q := hp{}
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && q[0].first < nums[i] {
			res[q[0].second] = nums[i]
			heap.Pop(&q)
		}
		for len(st) > 0 && nums[st[len(st) - 1]] < nums[i] {
			heap.Push(&q, Pair{nums[st[len(st) - 1]], st[len(st) - 1]})
			st = st[:len(st) - 1]
		}
		st = append(st, i)
	}
	return res
}