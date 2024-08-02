package main

import (
	"fmt"
)

func main() {
	var x string
    if x == "" {
        x = "default"
    }
    fmt.Println(x)
	fmt.Println(vowelStrings([]string{"are","amy","u"}, 0, 2))
}

func vowelStrings(words []string, left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		if s(words[i]) {
			ans++
		}
	}
	return ans
}

func s(str string) bool {
	last := str[len(str)-1]
	if str[0] == 'a' || str[0] == 'e' || str[0] == 'i' || str[0] == 'o' || str[0] == 'u' {
		if last == 'a' || last == 'e' || last == 'i' || last == 'o' || last == 'u' {
			return true
		}
	}
	return false
}