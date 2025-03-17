package main

import (
	"fmt"
)

func main() {
	fmt.Println(scoreOfString("hello"))
}

func scoreOfString(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i-1])- int(s[i]))
	}
	return ans
}

func abs (a int) int {
	if a < 0 {
		return -a
	}
	return a
}