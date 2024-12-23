package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nonSpecialCount(4, 16))
}

func nonSpecialCount(l int, r int) int {
	n := int(math.Sqrt(float64(r)))
	v := make([]int, n+1)
	res := r - l + 1
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			if i * i >= l && i * i <= r {
				res--
			}
			for j := i * 2; j <= n; j += i {
				v[j] = 1
			}
		}
	}
	return res
}