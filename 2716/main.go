package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimizedStringLength("aaabc"))
}

func minimizedStringLength(s string) int {
	m := map[rune]struct{}{}
	for _, ch := range s {
		m[ch] = struct{}{}
	}
	return len(m)
}