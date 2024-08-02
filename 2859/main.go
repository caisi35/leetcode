package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		s := strconv.FormatInt(int64(i), 2)
		t := 0
		for _, j := range s {
			if j == '1' {
				t++
			}
		}
		if t == k {
			ans += nums[i]
		}
	}
	return ans
}
