package main

import (
	"fmt"
	"sort"
)
func main() {
	fmt.Println(smallestRangeI([]int{1,3,6}, 3))	// 0
	fmt.Println(smallestRangeI([]int{0,10}, 2))		// 6

}

func smallestRangeI(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	sort.Ints(nums)
	r := nums[n-1] - nums[0]
	if r < 2 * k {
		return 0
	}
	ans := r - 2 * k
	return ans
}