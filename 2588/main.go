package main

import (
	"fmt"
)

func main() {
	fmt.Println(beautifulSubarrays([]int{1, 2, 3, 4, 5}))
}

func beautifulSubarrays(nums []int) int {
	count := 0
	xor := 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		xor ^= num
		if val, ok := m[xor]; ok {
			count += val
		}
		m[xor]++
	}
	return count
}
