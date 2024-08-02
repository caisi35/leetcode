package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))
}

func findTheArrayConcVal(nums []int) int64 {
	var ans int64
	n := len(nums)
	for i := len(nums) - 1; i >= len(nums)/2; i-- {
		if i == n-i-1 {
			ans += int64(nums[i])
			break
		}
		str := strconv.Itoa(nums[n-i-1]) + strconv.Itoa(nums[i])
		t, _ := strconv.Atoi(str)
		ans += int64(t)
	}
	return ans
}
