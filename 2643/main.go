package main

import (
	"fmt"
)

func main() {
	fmt.Print(rowAndMaximumOnes([][]int{{0,1}, {1,0}}))
}

func rowAndMaximumOnes(mat [][]int) []int {
	idx := 0
	mx := 0
	for i, nums := range mat {
		cnt := 0
		for _, n := range nums {
			if n == 1 {
				cnt++
			}
		}
		if cnt > mx {
			idx = i
			mx = cnt
		}
	}
	return []int{idx, mx}
}