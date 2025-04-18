package main

import "fmt"

func main() {
	fmt.Println(countBadPairs([]int{4,1,3,3}))	// 5
	fmt.Println(countBadPairs([]int{1,2,3,4,5}))	// 0
}

func countBadPairs(nums[]int) int64 {
	m := make(map[int]int)
	var res int64 = 0
	for i := 0; i < len(nums); i++ {
		k := nums[i] - i
		res += int64(i - m[k])
		m[k]++
	}
	return res
}

func countBadPairs2(nums[]int) int64 {
	ans := int64(0)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if i < j && nums[j] - nums[i] != j - i {
				ans++
			}
		}
	} 
	return int64(ans)
}