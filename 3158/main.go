package main

import "fmt"

func main() {
	fmt.Println(duplicateNumbersXOR([]int{1,2,1,3}))
	fmt.Println(duplicateNumbersXOR([]int{1,2,2, 1}))
}

func duplicateNumbersXOR(nums []int) int {
	ans := 0 
	m := make(map[int]int)
	for _, i := range nums {
		m[i]++
	}
	for k, v := range m {
		if v > 1 {
			ans ^= k
		}
	}
	return ans
}