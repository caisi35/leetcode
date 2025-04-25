package main

import "fmt"

func main() {
	fmt.Println(countInterestingSubarrays([]int{3,2,4}, 2, 1)) // 3
}

func countInterestingSubarrays(nums []int, modelo int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	var res int64 = 0
	var prefix int = 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		if nums[i] % modelo == k {
			prefix++
		}
		res += int64(cnt[(prefix - k + modelo) % modelo])
		cnt[prefix % modelo]++
	}
	return res
}