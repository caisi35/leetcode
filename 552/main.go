package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkRecord(2))	// 8
	fmt.Println(checkRecord(10101))	// 183236316

}

const mod = 1000000007
const mx = 100001
var memo [mx][2][3] int

func dfs(i, j, k int) int {
	if i == 0 {
		return 1
	}
	p := &memo[i][j][k]
	if *p > 0 {
		return *p
	}
	res := dfs(i-1, j, 0)
	if j == 0 {
		res += dfs(i-1, 1, 0)
	}
	if k < 2 {
		res += dfs(i-1, j, k+1)
	}
	*p = res % mod
	return *p
}

func checkRecord(n int) int {
	return dfs(n, 0, 0)
}