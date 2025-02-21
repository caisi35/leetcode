package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumWhiteTiles("10110101", 2, 2))		// 2
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	n := len(floor)
	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, numCarpets+1)
		for j := range d[i] {
			d[i][j] = 0x3f3f3f3f
		}
	}
	for j := 0; j <= numCarpets; j++ {
		d[0][j] = 0
	}
	for i := 1; i <= n ; i++ {
		d[i][0] = d[i-1][0]
		if floor[i-1] == '1' {
			d[i][0]++
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= numCarpets; j++ {
			d[i][j] = d[i-1][j]
			if floor[i-1] == '1' {
				d[i][j]++
			}
			d[i][j] = min(d[i][j], d[max(0, i - carpetLen)][j-1])
		}
	}
	return d[n][numCarpets]
}