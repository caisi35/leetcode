package main

import (
	"fmt"
)

func main() {
	// a, b  := 1.0, 4.0
	// fmt.Println(a | b)
	fmt.Println(maximumSumOfHeights([]int{3, 5, 3, 5, 1, 5, 4, 4, 4}))
}

func maximumSumOfHeights(maxHeights []int) int64 {
	max_index := 0
	res := 0
	n := len(maxHeights)
	for i := 1; i < n; i++ {
		if maxHeights[i] > maxHeights[max_index] {
			max_index = i
		}
	}

	for i := 0; i < n; i++ {
		if maxHeights[i] == maxHeights[max_index] {
			m := g(i, n , maxHeights)
			if m > res {
				res = m
			}
		}
	}

	return int64(res)
}

func g(max_index, n int, maxHeights []int ) int {
	res := 0

	b := make([]int, n)

	b[max_index] = maxHeights[max_index]
	v := 0
	if max_index+1 < n {
		v = maxHeights[max_index+1]
	}

	for i := max_index + 1; i < n; i++ {
		if maxHeights[i] > v {
			b[i] = v
		} else {
			b[i] = maxHeights[i]
			v = maxHeights[i]
		}
	}

	if max_index-1 >= 0 {
		v = maxHeights[max_index-1]
	}
	for i := max_index - 1; i >= 0; i-- {
		if maxHeights[i] > v {
			b[i] = v
		} else {
			b[i] = maxHeights[i]
			v = maxHeights[i]
		}
	}
	fmt.Println(b)
	for _, i := range b {
		res += i
	}
	return res
}