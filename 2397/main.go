package main

import (
	"fmt"
	"math/bits"
)

const (
	x = iota
	_
	y
	z = 9
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
	fmt.Println(maximumRows([][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 1},
	}, 2))
}

func maximumRows(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mask[i] += matrix[i][j] << (n - j - 1)
		}
	}
	res, limit := 0, 1<<n
	for cur := 1; cur < limit; cur++ {
		if bits.OnesCount(uint(cur)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < m; j++ {
			if (mask[j] & cur) == mask[j] {
				t++
			}
		}
		res = max(res, t)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
