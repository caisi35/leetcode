package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{1,2,3,4,7}, 3))
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1] - position[0]
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if check(position, m, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func check(position []int, m int, mid int) bool {
	c := 1
	prev := position[0]
	for i := 1; i < len(position); i++ {
		if position[i] - prev >= mid {
			c++
			prev = position[i]
		}
	}
	return c >= m
}