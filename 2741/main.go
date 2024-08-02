package main

import "fmt"

func main() {
	fmt.Println(specialPerm([]int{2,3,6}))
}

const mod = 1e9 + 7
func specialPerm(nums []int) int {
	n := len(nums)
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(state, i int) int {
		if f[state][i] != -1 {
			return f[state][i]
		}
		if state == (1 << i) {
			return 1
		}
		f[state][i] = 0
		for j := 0; j < n; j++ {
			if i == j || state >>j & 1 == 0 {
				continue
			}
			if nums[i] % nums[j] != 0 && nums[j] % nums[i] != 0 {
				continue
			}
			f[state][i] = (f[state][i] + dfs(state ^ (1 << i), j)) % mod
		}
		return f[state][i]
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + dfs((1 << n) - 1, i)) % mod
	}
	return res
}