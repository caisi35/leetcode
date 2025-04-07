package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsetXORSum([]int{5,1,6}))
}

func subsetXORSum(nums []int) int {
	return dfs(0, 0, nums)
}

func dfs(val int, idx int, nums []int) int {
	if idx == len(nums) {
		return val
	}
	return dfs(val ^ nums[idx], idx + 1, nums) + dfs(val, idx + 1, nums)
}