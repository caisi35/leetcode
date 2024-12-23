package main

import (
	"fmt"
)

func main() {
	fmt.Println(resultsArray([]int{3,2,3,2,3,2}, 2))
}

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	cnt := 0 
	ans := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		if i == 0 || nums[i] - nums[i-1] != 1 {
			cnt = 1
		} else {
			cnt++
		}
		if cnt >= k {
			ans[i-k+1] = nums[i]
		} else if i - k + 1 >= 0 {
			ans[i-k+1] = -1
		}
	}
	return ans
}