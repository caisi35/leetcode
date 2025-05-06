package main

import "fmt"

func main() {
	fmt.Println(numTilings(4))
}

func numTilings(n int) int {
	const mod int = 1e9+7
	dp := make([][4]int, n+1)
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (((dp[i-1][0]+dp[i-1][1])%mod+dp[i-1][2])%mod +dp[i-1][3]) % mod
	}
	return dp[n][3]
}