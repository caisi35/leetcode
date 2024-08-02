package main

import (
	"fmt"
)

func sum(x, y int) (t int, e error) {
	return x + y, nil
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(makeSmallestPalindrome("egcfe"))
}

func makeSmallestPalindrome(s string) string {
	n := len(s)
	res := make([]byte, n)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1]{
			if s[i] < s[n-i-1] {
				res[i] = s[i]
				res[n-i-1] = s[i]
			} else {
				res[i] = s[n-i-1]
				res[n-i-1] = s[n-i-1]
			}
		} else {
			res[i] = s[i]
			res[n-i-1] = s[n-i-1]
		}
	}
	if n % 2 != 0 {
		res[n / 2] = s[n/2]
	}
	return string(res)
}
