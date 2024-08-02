package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	m := make(map[int]int)
	for _, i := range candyType {
		m[i] = 1
		if len(m) == len(candyType) / 2 {
			return len(m)
		}
	}
	return len(m)
}
