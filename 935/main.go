package main

import (
	"fmt"
)

func main() {
	fmt.Println(knightDialer(3131))
}

func knightDialer(n int) int {
	const mod = 1_000_000_007
	moves := [][]int{
		{4,6},
		{6,8},
		{7,9},
		{4,8},
		{3,9,0},
		{},
		{1,7,0},
		{2,6},
		{1,3},
		{2,4},
	}
	d := [2][10]int{}
	for i := 0; i < 10; i++ {
		d[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		x := i & 1
		for j := 0; j < 10; j++ {
			d[x][j] = 0
			for _, k := range moves[j] {
				d[x][j] = (d[x][j] + d[x^1][k]) % mod
			}
		}
	}
	res := 0
	for _, x := range d[n%2] {
		res = (res + x) % mod
	}
	return res
}