package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{9,7,5,3}, 1))	// 4
	fmt.Println(minOperations([]int{2,1,2}, 2))		// -1
	fmt.Println(minOperations([]int{5,2,5,4,5}, 2))	// 2
}

func minOperations(nums []int, k int) int {

	m := map[int]bool{}
	for _, num := range nums {
		if num < k {
			return -1
		} else if num != k {
			m[num] = true
		}
	}

	return len(m)
}