package main

import "fmt"

func main() {
	fmt.Println(minCost([]int{20,1,15}, 5))
}

func sum(arr []int) int {
	ans := 0
	for _, a := range arr {
		ans += a
	}
	return ans
}

func minCost(nums []int, x int) int64 {
	n := len(nums)
	f := make([]int, n)
	copy(f, nums)
	ans := sum(f)
	for k := 1; k < n; k++ {
		for i := 0; i < n; i++ {
			f[i] = min(f[i], nums[(i + k) % n])
		}
		ans = min(ans, k * x + sum(f))
	}
	return int64(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}