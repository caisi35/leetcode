package main

import "fmt"

func main() {
	fmt.Println(longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(longestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
	fmt.Println(longestEqualSubarray([]int{1}, 1))
	fmt.Println(longestEqualSubarray([]int{2, 2}, 1))
	fmt.Println(longestEqualSubarray([]int{1, 2, 1}, 0))
	fmt.Println(longestEqualSubarray2([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(longestEqualSubarray2([]int{1, 1, 2, 2, 1, 1}, 2))
	fmt.Println(longestEqualSubarray2([]int{1}, 1))
	fmt.Println(longestEqualSubarray2([]int{2, 2}, 1))
	fmt.Println(longestEqualSubarray2([]int{1, 2, 1}, 0))
}

func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		c := k
		mx := 1
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				mx++
			} else if nums[i] != nums[j] && c > 0 {
				c--
			} else {
				break
			}
		}
		ans = max(mx, ans)
	}
	return ans
}

func longestEqualSubarray2(nums []int, k int) int {
	pos := make(map[int][]int)
	for i, num := range nums {
		pos[num] = append(pos[num], i)
	}

	ans := 0
	for _, vec := range pos {
		j := 0 
		for i := 0; i < len(vec); i++ {
			for vec[i] - vec[j] - (i - j) > k {
				j++
			}
			ans = max(ans, i - j + 1)
		}
	}
	return ans
}
