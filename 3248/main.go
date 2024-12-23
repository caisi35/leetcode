package main

import (
	"fmt"
)

func main() {
	fmt.Println(finalPositionOfSnake(3, []string{"DOWN", "RIGHT", "UP"}))
	fmt.Println(finalPositionOfSnake(2, []string{"RIGHT", "DOWN"}))
}

func finalPositionOfSnake(n int, commands []string) int {
	res := 0
	for _, c := range commands {
		if c == "RIGHT" {
			res++
		} else if c == "LEFT" {
			res--
		} else if c == "DOWN" {
			res += n
		} else {
			res -=n
		}
	}

	return res
}