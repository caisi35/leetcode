package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1,3,2,4,5}))
}

func countQuadruplets(nums []int) int64 {
	n := len(nums)
	pre := make([]int, n + 1)
	var ans int64 = 0
	for j := 0; j < n; j++ {
		suf := 0
		for k := n-1; k > j; k-- {
			if nums[j] > nums[k] {
				ans += int64(pre[nums[k]]) * int64(suf)
			} else {
				suf++
			}
		}
		for x := nums[j] + 1; x <= n; x++ {
			pre[x]++
		}
	}
	return ans
}