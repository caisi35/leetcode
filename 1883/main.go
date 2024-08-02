package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSkips([]int{1, 3, 2}, 4, 2))
	fmt.Println("foo:", foo())
}

func foo() int {
	var i int
	defer func() {
		fmt.Println("defer befroe i:", i)
	}()
	defer func() {
		i++
	}()
	return i
}

func minSkips(dist []int, speed int, hoursBefore int) int {
	const EPS = 1e-7
	n := len(dist)
	f := make([][]float64, n + 1)
	for i := range f {
		f[i] = make([]float64, n + 1)
		for j := range f[i] {
			f[i][j] = math.Inf(1)
		}
	}
	f[0][0] = 0.0
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j != i {
				f[i][j] = math.Min(f[i][j], math.Ceil(f[i-1][j] + float64(dist[i-1])/float64(speed)-EPS))
			}
			if j != 0 {
				f[i][j] = math.Min(f[i][j], f[i-1][j-1]+float64(dist[i-1])/float64(speed))
			}
		}
	}
	for j := 0; j <= n; j++ {
		if f[n][j] < float64(hoursBefore) + EPS {
			return j
		}
	}
	return -1
}