package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1,2,3,4,5}, []int{2,3,4,5,6}, 2))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	ans := int64(0)
	c1 := make(map[int]int)
	c2 := make(map[int]int)
	m := 0 
	for _, i := range nums1 {
		c1[i]++
		if i > m {
			m = i
		}
	}
	for _, j := range nums2 {
		c2[j]++
	}

	for a, v := range c2 {
		for b := a * k; b <= m; b += a * k {
			if _, ok := c1[b]; ok {
				ans += int64(v * c1[b])
			}
		}
	}
	return ans
}