package main

import "fmt"

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	var ans int64 = 0
	l := 0
	for l < n-1 {
		if nums[l] >= nums[l+1] {
			break
		}
		l++
	}
	if l == n-1 {
		return int64(n) * int64(n+1) / 2
	}
	ans += int64(l+2)
	for r := n-1; r > 0; r-- {
		if r < n-1 && nums[r] >= nums[r+1] {
			break
		}
		for l >= 0 && nums[l] >= nums[r] {
			l--
		}
		ans += int64(l+2)
	}
	return ans
}

func main() {
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
}
