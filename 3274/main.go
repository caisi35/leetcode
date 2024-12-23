package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkTwoChessboards("a1", "c3"))
	fmt.Println(checkTwoChessboards("a1", "h3"))
	fmt.Println(checkTwoChessboards("a7", "a6"))	// false
}

func checkTwoChessboards(coordinate1, coordinate2 string) bool {
	cs1 := coordinate1[0] - 'a' + 1 + coordinate1[1] - '0'
	cs2 := coordinate2[0] - 'a' + 1 + coordinate2[1] - '0'

	return cs1 % 2 == cs2 % 2
}