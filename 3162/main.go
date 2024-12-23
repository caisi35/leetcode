package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1,3,4}, []int{1,3,4}, 1))	// 5
	fmt.Println(numberOfPairs([]int{1,2,4,12}, []int{2,4}, 3))	// 2
}

func numberOfPairs(nums1 []int, nums2 []int, k int) (ans int) {
	for _, i := range nums1 {
		for _, j := range nums2 {
			if i % (j * k) == 0 {
				ans++
			}
		}
	}
	return
}