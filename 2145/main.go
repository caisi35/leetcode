package main

import "fmt"

func main() {
	fmt.Println(numberOfArrays([]int{1,-3,4}, 1, 6))	// 2
}

func numberOfArrays(differences []int, lower int, upper int) int {
	
	sum := 0
	i, j := 0, 0
	for _, num := range differences {
		sum += num
		i, j = min(i, sum), max(j, sum)
		if j - i > upper - lower {
			return 0
		}
	}
	return (upper - lower) - (j - i) +1
}