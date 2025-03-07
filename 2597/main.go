package main

import (
	"fmt"
)

func main() {
	fmt.Println(beautifulSubsets([]int{1,2,3,4,5}, 2))
}

func beautifulSubsets(nums []int, k int) int {
	count := 0
	n := len(nums)
	var backtrack func(index int, curr []int)
	backtrack = func(index int, curr []int) {
		if index == n {
			if len(curr) > 0 {
				count++
			}
			return
		}
		backtrack(index+1, curr)
		canAdd := true
		for _, num := range curr {
			if abs(nums[index]-num) == k {
				canAdd =false
				break
			}
		}
		if canAdd {
			curr = append(curr, nums[index])
			backtrack(index+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, []int{})
	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}