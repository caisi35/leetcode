package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumOfBeauties([]int{1,2,3}))
}

func sumOfBeauties(nums []int) int {
	n := len(nums)
	sum := 0
	pmax := make([]int, n)
	smin := make([]int, n)
	pmax[0] = nums[0]
	for i := 1; i < n; i++ {
		pmax[i] = max(pmax[i-1], nums[i])
	}
	smin[n-1] = nums[n-1]
	for i := n-2; i >= 0; i-- {
		smin[i] = min(smin[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if pmax[i-1] < nums[i] && nums[i] < smin[i+1] {
			sum += 2
			continue
		}
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			sum += 1
			continue
		}
	}
	return sum
}

func sumOfBeauties2(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 1; i < n-1; i++ {
		is := true
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				is = false
				break
			}
		}
		for k := i+1; k < n; k++ {
			if nums[k] <= nums[i] {
				is = false
				break
			}
		}
		if is {
			sum += 2
			continue
		}
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			sum += 1
			continue
		}
	}
	return sum
}