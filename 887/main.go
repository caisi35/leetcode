package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(3, 14))
}

func superEggDrop(k, n int) int {
	memo := make(map[int]int)
	var dp func(k,n int) int
	dp = func(k, n int) int {
		if _, ok := memo[n * 100 + k]; !ok {
			ans := 0
			if n == 0 {
				ans = 0
			} else if k == 1 {
				ans = n
			} else {
				lo, hi := 1, n
				for lo + 1 < hi {
					x := (lo + hi) / 2
					t1 := dp(k-1, x-1)
					t2 := dp(k, n-x)
					if t1 < t2 {
						lo = x
					} else if t1 > t2 {
						hi = x
					} else {
						hi = x
						lo = hi
					}
				}
				ans = 1 + min(max(dp(k-1, lo-1), dp(k, n-lo)), max(dp(k-1, hi-1), dp(k, n-hi)))
			}
			memo[n*100+k] = ans
		}
		return memo[n*100+k]
	}

	return dp(k, n)
}