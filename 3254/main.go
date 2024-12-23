package main

import (
	"fmt"
)

func main() {
	fmt.Println(resultsArray([]int{3,2,3,2,3,2}, 2))
}

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	for i := range res {
		res[i] = -1
	}
	for i := 0; i <= n - k; i++ {
		valid := true
		for j := i+1; j < i+k; j++ {
			if nums[j] - nums[j-1] != 1 {
				valid = false
				break
			}
		}
		if valid {
			res[i] = nums[i+k-1]
		}
	}
	return res
}