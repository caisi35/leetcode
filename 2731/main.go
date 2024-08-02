package main

import (
	"fmt"
	"sort"
)

func main() {
	m := []int{-2,0,2}
	fmt.Println(sumDistance(m, "RLL", 3))
}

func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	n := len(s)
	pos := make([]int, n)
	for i, v := range s {
		if v == 'R' {
			pos[i] = nums[i] +d
		} else {
			pos[i] = nums[i] - d
		}
	}
	sort.Ints(pos)
	res := 0
	for i := 1; i < n; i++ {
		res += (pos[i] - pos[i -1]) * i % mod * (n - i) % mod
		res %= mod
	}
	return res
}