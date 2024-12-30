package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isSubstringPresent("abcde")) 
	fmt.Println(isSubstringPresent("abcba")) 
}

func isSubstringPresent(s string) bool {
	for i := 1; i < len(s); i++ {
		if strings.Contains(s, string([]byte{s[i], s[i-1]})) {
			return true
		}

	}
	return false
}