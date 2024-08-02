package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumMountainRemovals([]int{2, 1, 1, 5, 6, 2, 3, 1}))
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	pre := getLISArray(nums)
	suf := getLISArray(reversed(nums))
	suf = reversed(suf)
	ans := 0
	for i := 0; i < n; i++ {
		if pre[i] > 1 && suf[i] > 1 {
			ans = max(ans, pre[i] + suf[i] - 1)
		}
	}
	return n- ans
}

func getLISArray(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	var seq []int
	for i := 0; i < n; i++ {
		it := sort.SearchInts(seq, nums[i])
		if it == len(seq) {
			seq = append(seq, nums[i])
			dp[i] = len(seq)
		} else {
			seq[it] = nums[i]
			dp[i] = it + 1
		}
	}
	return dp
}

func reversed(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[n - 1 - i]
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}