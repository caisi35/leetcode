package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countFairPairs([]int{0,1,7,4,4,5}, 3, 6))	// 6
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := int64(0)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r-l)
	}
	return ans
}

func countFairPairs2(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if lower <= nums[i] + nums[j] && nums[i] + nums[j] <= upper {
				ans++
			}
		}
	}
	return int64(ans)
}