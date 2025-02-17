package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSpecialInteger([]int{1,1}))
	fmt.Println(findSpecialInteger([]int{1,2,2,6,6,6,7,10}))
}

func findSpecialInteger(arr []int) int {
	ans := 0
	n := len(arr) 
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
		if float64(m[v]) / float64(n) > 0.25 {
			return v
		}
	}
	return ans
}