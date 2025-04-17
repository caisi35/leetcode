package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{3,1,2,2,2,1,3}, 2))
}

func countPairs(nums []int, k int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i+1; i < j && j < n; j++ {
			if nums[i] == nums[j] && i * j % k == 0 {
				ans++
			}
		}
	}
	return ans
}