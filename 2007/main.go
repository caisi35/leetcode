package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(f([]int{1, 3, 4, 2, 6, 8}))
}

func f(change []int) []int {
	sort.Ints(change)
	count := make(map[int]int)
	for _, num := range change {
		count[num]++
	}
	res := []int{}
	for _, a := range change {
		if count[a] == 0 {
			continue
		}
		count[a]--
		if count[a*2] == 0 {
			return []int{}
		}
		count[a*2]-- 
		res = append(res, a )
	}
	return res
}