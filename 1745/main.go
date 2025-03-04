package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPartitioning("abcasedfabab"))
}

func checkPartitioning(s string) bool {
	n := len(s)
	for i := 1; i < n-1; i++ {
		if isPalindrome(s[:i]) {
			for j := i + 1; j < n; j++ {
				if isPalindrome(s[i:j]) && isPalindrome(s[j:]) {
					return true
				}
			}
		}
	}
	return false
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n / 2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}