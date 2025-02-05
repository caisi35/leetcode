package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
	fmt.Println(sortArrayByParityII([]int{4,3}))
}

func sortArrayByParityII(nums []int) []int {
	ans := make([]int, len(nums))
	s := []int{}
	d := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 0 {
			s = append(s, nums[i])
		} else {
			d = append(d, nums[i])
		}
	}
	si := 0
	di := 0
	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			ans[i] = s[si]
			si++
		} else {
			ans[i] = d[di]
			di++
		}
	}
	return ans
}