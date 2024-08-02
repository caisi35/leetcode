package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // 1 12 13 4 5
	fmt.Println("a = ", a) // 1 12 13 4 5
	fmt.Println(minNonZeroProduct(2))
	fmt.Println(minNonZeroProduct(3))
	fmt.Println(minNonZeroProduct(4))
}

func minNonZeroProduct(p int) int {
	if p == 1 {
		return 1
	}
	mod := int64(1e9 + 7)
	x := fastPow(2, int64(p), mod) - 1
	y := int64(1) << uint(p-1)
	return int(fastPow(x-1, y-1, mod) * x % mod)
}

func fastPow(x, n, mod int64) int64 {
	res := int64(1)
	for n != 0 {
		if n&1 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n >>= 1
	}
	return res
}
