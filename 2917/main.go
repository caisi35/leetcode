package main

import (
	"fmt"
)

func findKOr(nums []int, k int) int {
	ans := 0 
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if (num >> i) & 1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			ans |= 1 << i
		}
	}
	return ans
}

type MyInt int

func (i MyInt) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i MyInt = 1
	i.PrintInt()
	fmt.Println(findKOr([]int{7, 12, 9, 8, 9, 15}, 4))
}
