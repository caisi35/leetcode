package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	h := 5
	w := 4
	horizontalCuts := []int{3, 1}
	verticalCuts := []int{1, 3}
	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts))
}

func max_slice(s []int, t int) int {
	ans, p := 0, 0
	for _, v := range s {
		ans = max(v-p, ans)
		p = v
	}
	return max(ans, t-p)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	ans := 0
	sort.Slice(horizontalCuts, func(i, j int) bool { return horizontalCuts[i] < horizontalCuts[j] })
	sort.Slice(verticalCuts, func(i, j int) bool { return verticalCuts[i] < verticalCuts[j] })
	// sort.Ints(horizontalCuts)
	// sort.Ints(verticalCuts)
	fmt.Println(horizontalCuts, verticalCuts)

	ans = max_slice(horizontalCuts, h) * max_slice(verticalCuts, w) % (7 + int(math.Pow10(9)))

	return ans
}
