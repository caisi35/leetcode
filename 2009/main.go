package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6, 6}))
}

func minOperations(nums []int) int {
	m := make(map[int]bool)
	n := []int{}
	for _, i := range nums {
		if ok := m[i]; !ok {
			m[i] = true
			n = append(n, i)
		}
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})
	// fmt.Println(n)
	ans := len(nums)
	j := 0 
	for i, l := range n {
		r := len(nums) + l - 1
		for j < len(n) && n[j] <= r {
			ans = min(ans, len(nums) - (j - i + 1))
			j++
		}
	}
	return ans
}
