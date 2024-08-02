package main

import "fmt"

func main() {
	list := make([]int, 0)

	list = append(list, 1)
	fmt.Println(list)
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
}

func maximumCount(nums []int) int {
	pos := 0
	neg := 0
	for _, num := range nums {
		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		}
	}
	return max(pos, neg)
}
