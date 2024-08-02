package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findValueOfPartition([]int{1,3,2,4}))
	fmt.Println(findValueOfPartition([]int{84,11,100,100,75}))

}

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ans := 100000000000
	for i := 0; i <= len(nums)-2; i++ {
		d := abs(nums[i], nums[i+1])
		ans = min(ans, d)
	}
	return ans
}

func abs(a, b int) int {
	t := a - b
	if t < 0 {
		return -t
	}
	return t
}