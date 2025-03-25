package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeValid("))()))", "010100"))
}

func canBeValid(s string, locked string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	b := 0 
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			b++
		} else {
			if s[i] == '(' {
				b++
			} else {
				b--
			}
		}
		if b < 0 {
			return false
		}
	}

	b = 0
	for i := len(s)-1; i >= 0; i-- {
		if locked[i] == '0' {
			b++
		} else {
			if s[i] == ')' {
				b++
			} else {
				b--
			}
		}
		if b < 0 {
			return false
		}
	}
	return true
}