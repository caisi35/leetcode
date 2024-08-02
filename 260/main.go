package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	ans := make(map[int]int)
	for _, i := range nums {
		if _, ok:= ans[i]; !ok {
			ans[i] = 1
		} else {
			ans[i] = 2
		}
	}
	// fmt.Println(ans)
	res := []int{}
	for k, v := range ans {
		if v == 1{
			res = append(res, k)
		}
	}
	return res
}
