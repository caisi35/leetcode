package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 3, 7))
	// fmt.Println(add([]int{1, 2}))
	fmt.Println(add([]int{1, 3, 7}...))
	fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	queue := make([]int, n)
	qi, qj := 0, 1
	for i := 1; i < n; i++ {
		for qi < qj && queue[qi] < i-k {
			qi++
		}
		dp[i] = dp[queue[qi]] + nums[i]
		for qi < qj && dp[queue[qj-1]] <= dp[i] {
			qj--
		}
		queue[qj] = i
		qj++
	}
	return dp[n-1]
}
