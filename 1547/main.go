package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minCost(9, []int{5,6,1,4,2}))
}

func minCost(n int, cuts []int) int {
	m := len(cuts)
	cuts = append([]int{0}, cuts...)
	cuts = append(cuts, n)
	sort.Ints(cuts)
	f := make([][]int, m+2)
	for i := range f {
		f[i] = make([]int, m+2)
	}
	for i := m; i >= 1; i-- {
		for j := i; j <= m; j++ {
			if i == j {
				f[i][j] = 0
			} else {
				f[i][j] = int(^uint(0) >> 1)
			}
			for k := i; k <= j; k++ {
				f[i][j] = min(f[i][j], f[i][k-1]+f[k+1][j])
			}
			f[i][j] += cuts[j+1] - cuts[i-1]
		}
	}
	return f[1][m]
}