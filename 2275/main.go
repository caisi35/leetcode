package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestCombination([]int{8,8}))
}

func largestCombination(candidates []int) int {
	m := func(k int) int {
		res := 0
		for _, num := range candidates {
			if num & (1 << k) != 0 {
				res++
			}
		}
		return res
	}
	res := 0 
	for i := 0; i < 24; i++ {
		res = max(res, m(i))
	}
	return res
}