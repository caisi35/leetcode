package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		result = append(result, append([]int(nil), path...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			backtrack(i+1, append(path, nums[i]))
		}
	}
	backtrack(0, []int{})
	return result
}

func main() {
	nums := []int{1,2,2}
	result := subsetsWithDup(nums)
	fmt.Println(result)
}