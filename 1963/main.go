package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSwaps("[][][][]"))
}

func minSwaps(s string) int {
	cnt, m := 0, 0 
	for _, ch := range s {
		if ch == '[' {
			cnt++
		} else {
			cnt--
			m = min(m, cnt)
		}
	}
	return (-m + 1) / 2
}