package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	// 0 -> 0
	// 1 -> 1
	// 2 -> 2
	// 3 -> 3

	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
}

func minimumSum(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	ans := -1
	for i := 0; i < len(nums)-2; i++ {
		for j := i; j < len(nums)-1; j++ {
			for k := j; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					sum := nums[i] + nums[j] + nums[k]
					if ans == -1 {
						ans = sum
					} else if sum < ans {
						ans = sum
					}
				}
			}
		}

	}
	return ans
}
