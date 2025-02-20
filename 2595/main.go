package main

import (
	"fmt"
)

func main() {
	fmt.Println(evenOddBit(66))
	fmt.Println(evenOddBit(17))
	fmt.Println(evenOddBit(50))
}

func evenOddBit(n int) []int {
	ans := make([]int, 2)
	even := 0
	odd := 0
	idx := 0
	for n > 0 {
		re := n % 2
		if re == 1 && idx%2 == 0 {
			even++
		} else if re == 1 && idx%2 == 1 {
			odd++
		}
		idx++
		n = n / 2
	}

	ans[0], ans[1] = even, odd
	return ans
}
