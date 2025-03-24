package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrefixes([]string{"a","a"}, "aa"))
}

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, word := range words {
		f := false
		for i := range word {
			if word > s {
				break
			}
			if word[i] != s[i] {
				f = false
				break
			}
			f = true
		}
		if f {
			ans++
		}
	}
	return ans
}