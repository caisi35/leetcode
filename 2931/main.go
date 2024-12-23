package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSpending([][]int{
		{8,5,2},
		{6,4,1},
		{9,7,3},
	}))
	fmt.Println(maxSpending([][]int{
		{10,8,6,4,2},
		{9,7,5,3,2},
	}))
}

func maxSpending(values [][]int) int64 {
	t := []int{}
	for _, v := range values {
		t = append(t, v...)
	}
	sort.Ints(t)
	ans := int64(0)
	for i := 0; i < len(t); i++ {
		ans += int64(t[i] * (i + 1))
	}
	return ans
}