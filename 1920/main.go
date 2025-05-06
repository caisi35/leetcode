package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{0,2,1,5,3,4}))		// 0 1, 2, 4, 5, 3
}

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}