package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{14,2,3,15,3,21}, 3))
}

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)
	mx := nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	pos := []int{-1}
	for i, v := range nums {
		if v == mx {
			pos = append(pos, i)
		}
	}
	left, right := 1, k
	var ans int64 = 0
	for right < len(pos) {
		ans += int64(pos[left]-pos[left-1]) * int64(n-pos[right])
		left++
		right++
	}
	return ans
}