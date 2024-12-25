package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(eatenApples([]int{1,2,3,5,2}, []int{3,2,1,4,2}))
}

func eatenApples(apples []int, days []int) int {
	h := hp{}
	ans := 0
	i := 0
	for ; i < len(apples); i++ {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if apples[i] > 0 {
			heap.Push(&h, pair{i+days[i], apples[i]})
		}
		if len(h) > 0 {
			h[0].left--
			if h[0].left == 0 {
				heap.Pop(&h)
			}
			ans++
		}
	}
	for len(h) > 0 {
		for len(h) > 0 && h[0].end <= i {
			heap.Pop(&h)
		}
		if len(h) == 0 {
			break
		}
		p := heap.Pop(&h).(pair)
		num := min(p.end-i, p.left)
		ans += num
		i += num
	}
	return ans
}

type pair struct {end, left int}
type hp []pair

func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].end < h[j].end}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(v interface{}) {*h = append(*h, v.(pair))}
func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v}