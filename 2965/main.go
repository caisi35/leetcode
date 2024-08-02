package main

import "fmt"

func main() {
	fmt.Println(findMissingAndRepeatedValues([][]int{
		{9,1,7},{8,9,2},{3,4,6},
	}))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	l := make([]int, len(grid) * len(grid) + 1)
	for _, g := range grid {
		for _, i := range g {
			l[i] += 1
		}
	}
	// fmt.Println(l)
	ans := make([]int, 2)
	for v, i := range l[1:] {
		if i == 2 {
			ans[0] = v +1
		} else if i == 0 {
			ans[1] = v + 1
		}
	}
	return ans
}