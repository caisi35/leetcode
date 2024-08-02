package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1,2,3,2,1}))
}

func findPeakElement(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := 1
	right, left := len(nums)-2, 1
	for left <= right {
		if nums[left] > nums[left-1] && nums[left] > nums[left+1] {
			return left
		}
		if nums[right] > nums[right-1] && nums[right] > nums[right+1] {
			return right
		}
		left++
		right--
	}
	if nums[right] > nums[left] {
		res = len(nums)-2
	}
	if nums[1] < nums[0] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums)-1
	}
	return res
}
