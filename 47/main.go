package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))

	sort.Ints(nums)
	backtrack(nums, path, used, &result)
	return result
}

func backtrack(nums, path []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack(nums, path, used, result)
		used[i] = false
		path = path[:len(path)-1]
	}
}