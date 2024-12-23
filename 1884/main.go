package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoEggDrop(100))
}

func twoEggDrop(n int) int {
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt32 / 2
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++{
			f[i] = min(f[i], max(k-1, f[i-k]) + 1)
		}
	}
	return f[n]
}