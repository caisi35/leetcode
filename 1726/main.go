package main

import (
	"fmt"
)

func tupleSameProduct(nums []int) int {
	ans := 0
	n := len(nums)
	t := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			t[nums[i]*nums[j]]++
		}
	}
	fmt.Println(t)
	for _, v := range t {
		ans += v * (v - 1) * 4
	}

	return ans
}

func main() {
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10}))
}
