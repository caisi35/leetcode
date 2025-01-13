package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(generateKey(1, 10, 1000))		// 0
	fmt.Println(generateKey(987, 879, 798))	 	// 777
	fmt.Println(generateKey(1140, 1851, 2057))	// 1040
}

func generateKey(num1 int, num2 int, num3 int) int {
	ans := ""
	nums1 := get(num1)
	nums2 := get(num2)
	nums3 := get(num3)
	// fmt.Println(nums1, nums2, nums3)
	a := []int{}
	for i := 0; i < 4; i++ {
		m := min(nums1[i], nums2[i],nums3[i])
		a = append(a, m)
	}
	fmt.Println(a)
	for i := range a {
		ans += fmt.Sprintf("%d", a[i])
	}
	fmt.Println(ans)
	r, _ := strconv.Atoi(ans) 
	return r
}

func get(num int) []int {
	nums := make([]int, 4)
	i := 3
	for num > 0 {
		m := num % 10
		nums[i] = m
		if m != 0 {
			num -= m
			num /= 10
		} else {
			num /= 10
		}
		i--
	}
	
	return nums
}