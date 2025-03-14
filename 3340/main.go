package main

import (
	"fmt"
)

func main() {
	fmt.Println(isBalanced("1234"))
	fmt.Println(isBalanced("2222"))
}

func isBalanced(num string) bool {
	o := 0
	j := 0
	for i, v := range num {
		if i % 2 == 0 {
			o += int(v - '0')
		} else {
			j += int(v - '0')
		}
	}
	return o == j
}