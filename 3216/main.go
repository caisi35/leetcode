package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSmallestString("001"))
	fmt.Println(getSmallestString("45320"))
}

func getSmallestString(s string) string {
	ans := []byte{}
	f := false
	for i := 1; i < len(s); i++ {
		// fmt.Println(s[i]-48)
		num1 := s[i-1]-48
		num2 := s[i]-48
		if num1 % 2 == num2 % 2 && num1 > num2 && !f {
			ans = append(ans, s[i], s[i-1])
			ans = append(ans, s[i+1:]...)
			f = true
			break
		} else {
			ans = append(ans, s[i-1])
		}
	}
	if !f {
		return s
	}
	return string(ans)
}