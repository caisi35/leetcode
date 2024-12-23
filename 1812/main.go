package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareIsWhite("a1"))
	fmt.Println(squareIsWhite("h3"))
	fmt.Println(squareIsWhite("c7"))
}

func squareIsWhite(coordinates string) bool {
	c1 := 'a' - coordinates[0] + 1
	c2 := coordinates[1]
	return (c1 + c2) % 2 != 0 
}