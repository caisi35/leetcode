package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{2,5,6,0,0,1,2}, 0))
}

func search(nums []int, target int) bool {
	for i := range nums {
		if nums[i] == target {
			return true
		}
	}
	return false
}