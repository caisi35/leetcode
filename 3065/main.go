package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([]int{2,11,10,1,3}, 10))
}

func minOperations(nums []int, k int) int {
	ans := 0 
	for i := range nums {
		if nums[i] < k {
			ans++
		}
	}
	return ans
}