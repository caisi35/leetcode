package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSum(5, 4))
}

func minimumSum(n int, k int) int {
	sum := 0
	skip := make(map[int]bool)
	num := 1
	for i := 0; i < n; {
		if !skip[num] {
			sum += num
			if num < k {
				skip[k-num] = true
			}
			i++
		}
		num++
	}
	return sum
}