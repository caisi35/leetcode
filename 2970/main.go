package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])

	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
	fmt.Println(incremovableSubarrayCount([]int{6, 5, 7, 8}))
}

func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isIncreasing(nums, i, j) {
				res++
			}
		}
	}
	return res
}

func isIncreasing(nums []int, l, r int) bool {
	for i := 1; i < len(nums); i++ {
		if i >= l && i <= r+1 {
			continue
		}
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	if l-1 >= 0 && r+1 < len(nums) && nums[r+1] <= nums[l-1] {
		return false
	}
	return true
}
