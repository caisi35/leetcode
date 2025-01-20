package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxValues([]int{2,6,7}, 1)) 	// 5
	fmt.Println(maxValues([]int{4,2,5,6,7}, 2))	// 2
}

func maxValues(nums []int, k int) int {
	findORs := func(nums []int, k int) []map[int]bool {
		dp := make([]map[int]bool, 0)
		prev := make([]map[int]bool, k+1)
		for i := 0; i <= k; i++ {
			prev[i] = make(map[int]bool)
		}
		prev[0][0] = true
		for i := 0; i < len(nums); i++ {
			for j := min(k-1, i+1); j >= 0; j-- {
				for x := range prev[j] {
					prev[j+1][x | nums[i]] = true
				}
			}
			cur := make(map[int]bool)
			for key := range prev[k] {
				cur[key] = true
			}
			dp = append(dp, cur)
		}
		return dp
	}
	A := findORs(nums, k)
	reverse(nums)
	B := findORs(nums, k)
	mx := 0
	for i := k-1; i < len(nums)-k; i++ {
		for a := range A[i] {
			for b := range B[len(nums)-i-2] {
				if a ^ b > mx {
					mx = a ^ b
				}
			}
		}
	}
	return mx
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}