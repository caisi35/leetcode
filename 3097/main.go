package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumSubarrayLength([]int{1,2,3}, 2))
}

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	bits := make([]int, 30)
	res := math.MaxInt32

	for left, right := 0, 0; right < n; right++ {
		for i := 0; i < 30; i++ {
			bits[i] += (nums[right] >> i) & 1
		}
		for left <= right && calc(bits) >= k {
			res = min(res, right - left + 1)
			for i := 0; i < 30; i++ {
				bits[i] -= (nums[left] >> i) & 1
			}
			left++
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func calc(bits []int) int {
	ans := 0
	for i := 0; i < len(bits); i++ {
		if bits[i] > 0 {
			ans |= 1 << i
		}
	}
	return ans
}