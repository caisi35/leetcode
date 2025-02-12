package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSize([]int{9}, 2))
}

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, 0
	for num := range nums {
		if nums[num] > r {
			r = nums[num]
		}
	}

	for l < r {
		mid := l + (r - l) / 2
		if canAchieve(nums, maxOperations, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func canAchieve(nums []int, maxOperations int, maxCost int) bool {
	operations := 0
	for _, num := range nums {
		operations += (num-1) / maxCost
		if operations > maxOperations {
			return false
		}
	}
	return operations <= maxOperations
}