package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0,1,0,0,1}))
	fmt.Println(numberOfAlternatingGroups([]int{0,0,1}))
}

func numberOfAlternatingGroups(colors []int) int {
	ans := 0
	n := len(colors)
	for i := 0; i < n; i++ {
		if colors[i] != colors[(i-1+n)%n] && colors[i] != colors[(i+1)%n] {
			ans += 1
		}
	}
	return ans
}