package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56}
	r := Constructor(arr)
	fmt.Println(r.Query(1, 2, 4))		// 1
	fmt.Println(r.Query(0, 11, 33)) 	// 2
}

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := map[int][]int{}
	for i, v := range arr {
		m[v] = append(m[v], i)
	}
	return RangeFreqQuery{m}
}

func (r *RangeFreqQuery) Query(left int, right int, value int) int {
	ans := 0
	var arr []int
	arr, ok := r.m[value]
	if !ok {
		return ans
	}
	l := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= left
	})
	ri := sort.Search(len(arr), func(i int) bool {
		return arr[i] > right
	})
	return ri - l
}
