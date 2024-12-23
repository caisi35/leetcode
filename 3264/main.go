package main

import (
	"fmt"
)

func main() {
	fmt.Println(getFinalState([]int{2,1,3,5,6}, 5, 2))
	fmt.Println(getFinalState([]int{7,71,15}, 10, 3))	// [567, 639, 1215]
}

func getFinalState(nums []int, k int, multipliter int) []int {
	for i := 0; i < k; i++ {
		idx := get_min(nums)
		nums[idx] = nums[idx] * multipliter
	}
	return nums
}

func get_min(nums []int) int {
	m := 100000000
	idx := 0
	for i := range nums {
		if nums[i] < m {
			m = nums[i]
			idx = i
		}
	}
	return idx
}