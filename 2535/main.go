package main 

import (
	"fmt"
	"strconv"
)

func main () {
	fmt.Println(differenceOfSum([]int{1,2,3,11}))
}

func differenceOfSum(nums []int) int {
	sum := 0
	num_s := 0
	for _, num := range nums {
		sum += num
		s := strconv.Itoa(num)
		for _, ch := range s {
			num_s += int(ch - '0')
		}
	}
	ans := sum - num_s
	if ans < 0 {
		return -ans
	}
	return ans
}