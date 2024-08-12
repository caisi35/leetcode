package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10}))	// -2
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))	// 2
	fmt.Println(minimumAddedInteger([]int{4,6,3,1,4,2,10,9,5}, []int{5,10,3,2,6,1,9})) // 0
	fmt.Println(minimumAddedInteger([]int{10,2,8,7,5,6,7,10}, []int{5,8,5,3,8,4})) // -2

}

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 2; i >= 0; i-- {
		left, right := i + 1, 1
		for left < len(nums1) && right < len(nums2) {
			if nums1[left] - nums2[right] == nums1[i] - nums2[0] {
				right++
			}
			left++
		}
		if right == len(nums2) {
			return nums2[0] - nums1[i]
		}
	}
	return 0
}
