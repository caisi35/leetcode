package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3}))	// 5
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,2,3}))	// 7
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	first, second := 2, 2
	for second < n {
		if nums[first-2] != nums[second] {
			nums[first] = nums[second]
			first++
		}
		second++
	}
	return first
}