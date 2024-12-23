package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAverage([]int{7,8,3,4,15,13,4,1}))
}

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	ans := 10000.0
	for i := 0; i < len(nums) / 2; i++ {
		avg := float64(nums[i] + nums[len(nums)-i-1]) / 2.0
		if avg < ans {
			ans = avg
		}
	}
	return ans
}