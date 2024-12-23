package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumOfMarkeIndices([]int{1, 2, 3, 4}))
}

func maxNumOfMarkeIndices(nums []int) int {
	sort.Ints(nums)
	ans := 0 
	n := len(nums)
	m := n / 2
	for i, j := 0, m; i < m && j < n; i++ {
		for ; j < n && 2 * nums[i] > nums[j]; {
			j++
		}
		if j < n {
			ans += 2
			j++
		}
	}
	return ans
}