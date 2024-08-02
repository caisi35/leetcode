package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
	fmt.Println(alternatingSubarray2([]int{21, 9, 5}))
}

func alternatingSubarray(nums []int) int {
	ans := -1

	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l := j - i + 1
			if nums[j]-nums[i] == (l-1)%2 {
				if l > ans {
					ans = l
				}
			} else {
				break
			}
		}
	}
	return ans
}

func alternatingSubarray2(nums []int) int {
	res := -1
	n := len(nums)
	firstIndex := 0
	for i := 1; i < n; i++ {
		length := i - firstIndex + 1
		if nums[i] - nums[firstIndex] == (length - 1) % 2 {
			if length > res {
				res = length
			}
		} else {
			if nums[i] - nums[i - 1] == 1 {
				firstIndex = i - 1
				if res < 2 {
					res = 2
				}
			} else {
				firstIndex = i
			}
		}
	}
	 return res
}