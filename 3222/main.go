package main

import (
	"fmt"
)

func main() {
	fmt.Println(losingPlayer(2, 7))
	fmt.Println(losingPlayer(4, 11))
}

func losingPlayer(x int, y int) string {
	f := false
	for x >= 1 && y >= 4 {
		f = !f
		x--
		y -= 4
	}
	if f {
		return "Alice"
	}
	return "Bob"
}