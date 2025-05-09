package main

import "fmt"

func main() {
	fmt.Println(countGoodNumbers(100))
}

func countGoodNumbers(n int64) int {
	mod := int64(1e9 + 7)

	quickmul := func(x, y int64) int64 {
		ret := int64(1)
		mul := x
		for y > 0 {
			if y % 2 == 1 {
				ret = ret * mul % mod
			}
			mul = mul * mul % mod
			y /= 2
		}
		return ret
	}
	return int(quickmul(5, (n+1)/2) * quickmul(4, n / 2) % mod)
}