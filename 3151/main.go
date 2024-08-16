package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(isArraySpecial([]int{4,1,2,3}))
	fmt.Fprintln(os.Stderr, " + ")
}

func isArraySpecial(nums []int) bool {
	ans := true 
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums)-1; i++ {
		if (nums[i] + nums[i+1]) % 2 == 0 {
			return false
		}
	}
	return ans
}