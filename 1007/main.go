package main

import "fmt"

func main() {
	fmt.Println(minDominoRotations([]int{1,2,3,4}, []int{1,2,3,4}))
}

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	r := check(tops[0], tops, bottoms, n)
	if r != -1 || tops[0] == bottoms[0] {
		return r
	}
	return check(bottoms[0], tops, bottoms, n)
}

func check(x int, tops, bottoms []int, n int) int {
	ra, rb := 0, 0
	for i := 0; i < n; i++ {
		if tops[i] != x && bottoms[i] != x {
			return -1
		} else if tops[i] != x {
			ra++
		} else if bottoms[i] != x {
			rb++
		}
	}
	return min(ra, rb)
}