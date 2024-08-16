package main

import (
	"fmt"
)

func main() {
	fmt.Println(isArraySpecial([]int{3,4,1,2,6, 1}, [][]int{{0,4}}))	// false
	fmt.Println(isArraySpecial([]int{4,3,1,6}, [][]int{{0,2}, {2,3}}))	// false true
	fmt.Println(isArraySpecial([]int{1}, [][]int{{0,0}}))	// true
	fmt.Println(isArraySpecial([]int{1, 8}, [][]int{{1, 1}}))	// true
	fmt.Println(isArraySpecial([]int{2,7,2}, [][]int{{0,0},{1, 2}}))	// true, true

}

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		if (nums[i] ^ nums[i-1]) & 1 == 1 {
			dp[i] = dp[i-1] + 1
		}
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		res[i] = dp[y] >= y - x + 1
	}
	return res
}

// 超时
func isArraySpecial2(nums []int, queries [][]int) []bool {

	sub := [][]int{}
	first, last := 0, 0
	for i := 1; i < len(nums); i++ {
		if (nums[i-1]+nums[i]) % 2 != 0 {
			last += 1
			if i == len(nums)-1 {
				sub = append(sub, []int{first, last})
			}
		} else {
			sub = append(sub, []int{first, last})
			first = i
			last = i
		}
	}
	// fmt.Println(sub)

	ans := make([]bool, len(queries))

	for i, q := range queries {
		if q[0] == q[1] && q[0] < len(nums) {
			ans[i] = true
			continue
		}
		for _,s := range sub {
			if q[0] >= s[0] && q[1] <= s[1] {
				ans[i] = true
			}
		}
	}
	return ans
}