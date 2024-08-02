package main

import (
	"fmt"
)

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": &Math{2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
