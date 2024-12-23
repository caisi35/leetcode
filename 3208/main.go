package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{1,1,0,1}, 4))
}

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	res, cnt := 0, 1
	for i := -k + 2; i < n; i++ {
		if colors[(i+n) % n] != colors[(i-1+n)%n] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			res++
		}
	}
	return res
}