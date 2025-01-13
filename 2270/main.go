package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToSplitArray([]int{10,4,-8,7}))
	fmt.Println(waysToSplitArray([]int{2,3,1,0}))
}

func waysToSplitArray(nums []int) int {
	res := 0
	l, r := 0, 0
	for i := range nums {
		r += nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		l += nums[i]
		r -= nums[i]
		if l >= r {
			res++
		}
	}
	return res
}