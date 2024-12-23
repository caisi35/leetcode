package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNonDuplicate([]int{1,1,2,2,3,4,4,5,5,6,6}))
	fmt.Println(singleNonDuplicate([]int{1,1,2,2,3,3,4,4,5,5,6,6,7}))
	fmt.Println(singleNonDuplicate([]int{1,1,2,2,3,3,4,4,5,5,6}))
	fmt.Println(singleNonDuplicate([]int{1,2,2,3,3,4,4,5,5,6}))
	fmt.Println(singleNonDuplicate([]int{1,1,2,3,3}))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)-1
	for i := 0; i <= len(nums) / 2; i+=2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		if nums[n-i] != nums[n-i-1] {
			return nums[n-i]
		}
	}
	return 0
}