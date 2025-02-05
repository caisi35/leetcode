package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("abca"))                   // true
	fmt.Println(validPalindrome("abc"))                    // false
	fmt.Println(validPalindrome("deeee"))                  // true
	fmt.Println(validPalindrome("tebbem"))                 // false
	fmt.Println(validPalindrome("eeccccbebaeeabebccceea")) // false
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))  // true
}

func validPalindrome(s string) bool {
	ans := true
	first := 0
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[first] == s[last] {
			first++
			last--
		} else {
			f1, f2 := true, true
			for i, j := first, last-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					f1 = false
					break
				}
			}
			for i, j := first+1, last; i < j; i, j= i+1, j-1 {
				if s[i] != s[j] {
					f2 = false
					break
				}
			}
			return f1 || f2
		}
	}

	return ans
}
