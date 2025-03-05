package main

import (
	"fmt"
)

func main() {
	fmt.Println(breakPalindrome("aba"))
}

func breakPalindrome(palindrome string) string {
	if len(palindrome) <= 1{
		return ""
	}

	s := []rune(palindrome)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != 'a' {
			s[i] = 'a'
			return string(s)
		}
	}
	s[len(s)-1] = 'b'
	return string(s)
}