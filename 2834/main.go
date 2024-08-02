package main

import (
	"fmt"
)

func minimumPossibleSum(n int, target int) int {
	const mod = 1000000007
	m := target / 2
	if n <= m {
		return ((1 + n) * n / 2) % mod
	}
	return (((1 + m) * m / 2) + ((target + target + (n - m) - 1) * (n - m) / 2)) % mod
}

func main() {
	fmt.Println(minimumPossibleSum(2, 3))
}
