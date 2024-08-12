package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(addedInteger([]int{2,6,4}, []int{9,7,5}))
}

func addedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    return -(nums1[0] - nums2[0])
}