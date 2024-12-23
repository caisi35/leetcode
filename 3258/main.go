package main

import (
	"fmt"
)

func main() {
	fmt.Println(countKConstraintSubstrings("1010101", 2))
}

func countKConstraintSubstrings(s string, k int) int {
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		count := [2]int{}
		for j := i; j < n; j++ {
			count[int(s[j]-'0')]++
			if count[0] > k && count[1] > k {
				break
			}
			res++
		}
	}
	return res
}