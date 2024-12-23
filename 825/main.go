package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numFriendRequests([]int{20,30,100,110,120}))	// 3
	fmt.Println(numFriendRequests([]int{16,17,18}))	// 2
	fmt.Println(numFriendRequests([]int{16,16}))	// 2
}

func numFriendRequests(ages []int) int {
	left, right := 0, 0
	ans := 0
	sort.Ints(ages)
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left] * 2 <= age + 14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		ans += right-left
	}
	return ans
}