package main

import (
	"fmt"
)

type MyInt int

func (i MyInt) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i MyInt = 1
	i.PrintInt()
	fmt.Println(maximizeSum([]int{1, 2, 3, 4, 5}, 3))
}

func maximizeSum(nums []int, k int) int {
	ans := 0
	m := 0
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	// ans += m
	for k > 0 {
		ans += m
		k -= 1
		m += 1
	}
	return ans
}

