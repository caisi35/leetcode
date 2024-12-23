package main

import (
	"fmt"
)

func main() {
	fmt.Println(semiOrderedPermutation([]int{2,4,1,3}))
}

func semiOrderedPermutation(nums []int) int {
	last, first := 0, 0
	n := len(nums)
	for i, v := range nums {
		if v == 1 {
			first = i
		}
		if v == n {
			last = i
		}
	}
	if last < first {
		return first + n - 1 - last - 1
	}
	return first + n - 1 - last
}