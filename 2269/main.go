package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(divisorSubstrings(240, 2)) // 2
}

func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	ans := 0
	for i := 0; i <= len(s)-k; i++ {
		sub := s[i:i+k]
		subn, _ := strconv.Atoi(sub)
		if subn != 0 && num % subn == 0 {
			ans++
		}
	}
	return ans
}