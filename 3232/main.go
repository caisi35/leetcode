package main

import (
	"fmt"
)

func main() {
	fmt.Println(canAliceWin([]int{5,5,5,25}))
	fmt.Println(canAliceWin([]int{1,2,3,4,10}))
	fmt.Println(canAliceWin([]int{1,2,3,4,5,14}))
}

func canAliceWin(nums []int) bool {
	sum1, sum2 := 0, 0
	for _, v := range nums {
		if v < 10 {
			sum1 += v
		} else {
			sum2 += v
		}
	}
	return sum1 != sum2
}