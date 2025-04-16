package main

import "fmt"

func main() {
	fmt.Println(countGood([]int{3,1,4,3,2,2,4}, 2))	// 4
}

func countGood(nums []int, k int) int64 {
	n := len(nums)
	same, right := 0, -1
	cnt := make(map[int]int)
	var ans int64 = 0
	for left := 0; left < n; left++ {
		for same < k && right + 1 < n {
			right++
			same += cnt[nums[right]]
			cnt[nums[right]]++
		}
		if same >= k {
			ans += int64(n-right)
		}
		cnt[nums[left]]--
		same -= cnt[nums[left]]
	}
	return ans
}