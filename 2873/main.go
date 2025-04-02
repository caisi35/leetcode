package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumTripletValue([]int{12,6,1,2,7})) 	// 77
	fmt.Println(maximumTripletValue([]int{2,3,1})) // 0
}

func maximumTripletValue(nums []int) int64 {
	ans := int64(0)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			for k := j+1; k < n; k++ {
				ans = max(ans, int64((nums[i]-nums[j]) * nums[k]))
			}
		}
	}
	return ans
}