package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("abasefasacacaaa"))
}

func partition(s string) [][]string {
	var r [][]string
	var c []string
	backtrack(s, 0, c, &r)
	return r
}

func backtrack(s string, start int, curr []string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			curr = append(curr, s[start:i+1])
			backtrack(s, i+1, curr, res)
			curr = curr[:len(curr)-1]
		}
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}