package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMatrix([]int{1,3,4,1,2,3,1}))
}

func findMatrix(nums []int) [][]int {
	ans := [][]int{}
	m := map[int]int{}
	mx := 0
	for _, num := range nums {
		m[num]++
		if m[num] > mx {
			mx = m[num]
		}
	}
	for i := 0; i < mx; i++ {
		ans = append(ans, []int{})
	}
	for k, v := range m {
		for i := 0; i < v; i++ {
			ans[i] = append(ans[i], k)
		}
	}
	return ans
}

func in(num int, nums []int) bool {
	for _, i := range nums {
		if i == num {
			return true
		}
	}
	return false
}