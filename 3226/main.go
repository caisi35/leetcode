package main

import (
	"fmt"
)

func main() {
	fmt.Println(minChanges(14, 13))
}

func minChanges(n int, k int) int {
	var res int 
	for n > 0 || k > 0 {
		if (n & 1) == 0 && (k & 1) == 1 {
			return -1
		}
		if (n & 1) == 1 && (k & 1) == 0 {
			res++
		}
		n >>= 1
		k >>= 1
	}
	return res
}