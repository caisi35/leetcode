package main

import "fmt"

func main() {
	fmt.Println(maxArrayValue([]int{2, 3, 7, 9, 3}))
	fmt.Println(maxArrayValue([]int{5, 3, 3}))
	fmt.Println(maxArrayValue([]int{91, 50}))
}

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	ans := nums[n-1]
	for i := n - 1; i > 0; i-- {
		if nums[i] >= nums[i-1] {
			nums[i-1] = nums[i] + nums[i-1]
		}
		if nums[i-1] > ans {
			ans = nums[i-1]
		}
	}
	return int64(ans)
}
