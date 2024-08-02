package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
	fmt.Println(paintWalls([]int{2, 3, 4, 2}, []int{1, 1, 1, 1}))
	fmt.Println(paintWalls([]int{26, 53, 10, 24, 25, 20, 63, 51}, []int{1, 1, 1, 1, 2, 2, 2, 1}))

}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, 2 * n + 1)
	for k := range f {
		f[k] = math.MaxInt / 2
	}

	f[n] = 0
	for i := 0; i < n; i++ {
		g := make([]int, n * 2 + 1)
		for k := range g {
			g[k] = math.MaxInt / 2
		}
		for j := 0; j <= n * 2; j++ {
			g[min(j + time[i], n * 2)] = min(g[min(j + time[i], n * 2)], f[j] + cost[i])
			if j > 0 {
				g[j - 1] = min(g[j - 1], f[j])
			}
		}
		f = g
	}
	res := math.MaxInt
	for i := n; i < len(f); i++ {
		res = min(res, f[i])
	}
	return res
}
