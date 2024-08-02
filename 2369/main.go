package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:] // [2, 3]
	fmt.Println(s2)
	s2[1] = 4 // [2, 4]
	fmt.Println(s2)
	fmt.Println(s1)          // [1, 2, 4]
	s2 = append(s2, 5, 6, 7) // [2, 4, 5, 6, 7]
	fmt.Println(s1)          // [1, 2, 4]
	fmt.Println(s2)

	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))
}

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		if i >= 2 {
			dp[i] = dp[i-2] && validTwo(nums[i-2], nums[i-1])
		}
		if i >= 3 {
			dp[i] = dp[i] || (dp[i-3] && validThree(nums[i-3], nums[i-2], nums[i-1]))
		}
	}
	return dp[n]
}

func validTwo(n1, n2 int) bool {
	return n1 == n2
}

func validThree(n1, n2, n3 int) bool {
	return (n1 == n2 && n1 == n3) || (n1+1 == n2 && n2+1 == n3)
}
