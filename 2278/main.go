package main

import (
	"fmt"
)

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))	// 33
	fmt.Println(percentageLetter("jjjj", 'k'))	// 33
	fmt.Println(percentageLetter("vmvvvvvzrvvpvdvvvvyfvdvvvvpkvvbvvkvvfkvvvkvbvvnvvomvzvvvdvvvkvvvvvvvvvlvcvilaqvvhoevvlmvhvkvtgwfvvzy", 'v'))	// 59
}

func percentageLetter(s string, letter byte) int {
	c := 0
	for _, ch := range s {
		if byte(ch) == letter {
			c++
		}
	}
	return int(100 * float32(c) / float32(len(s)))
}