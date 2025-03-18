package main

import (
	"fmt"
)

func main() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func diagonalPrime(nums [][]int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
		if isPrime(nums[i][len(nums[i])-i-1]) {
			ans = max(ans, nums[i][len(nums)-i-1])
		}
	}
	return ans
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
