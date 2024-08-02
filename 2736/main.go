package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{
		{4, 1}, {1, 3}, {2, 5},
	}))
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, querie := range queries {
		q1, q2 := querie[0], querie[1]
		m := -1
		for i := 0; i < len(nums1); i++ {
			s :=  nums1[i] + nums2[i]
			if nums1[i] >= q1 && nums2[i] >= q2 && s > m {
				m = s
			}
		}
		ans[i] = m
	}
	return ans
}
