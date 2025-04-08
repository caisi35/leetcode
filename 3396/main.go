package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{1,2,3,4,2,3,3,5,7}))	// 2
	fmt.Println(minimumOperations([]int{4,5,6,4,4}))	// 2
	fmt.Println(minimumOperations([]int{6,7,8,9}))	// 0
}

func minimumOperations(nums []int) int {
	n := len(nums)
	m := map[int]int{}
	idx := -1
	for i := n-1; i >= 0; i-- {
		if _, ok := m[nums[i]]; ok {
			idx = i+1
			break
		} else {
			m[nums[i]] = i
		}
	}
	ans := 0
	if idx > 0 {
		fmt.Println(idx, idx % 3)
		ans = idx / 3
		if idx % 3 != 0 {
			ans++
		}
	}
	return ans
}