package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1) // 1 2 3 4 5
	fmt.Println(maximumBinaryString("000110"))
	fmt.Println(maximumBinaryString2("000110"))

}

func maximumBinaryString(binary string) string {
	n := len(binary)
	s := []rune(binary)
	j := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			for j <= i || (j < n && s[j] == '1') {
				j++
			}
			if j < n {
				s[j] = '1'
				s[i] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}

func maximumBinaryString2(binary string) string {
	n := len(binary)
	i := strings.Index(binary, "0")
	if i < 0 {
		return binary
	}
	zeros := strings.Count(binary, "0")
	s := []rune(strings.Repeat("1", n))
	s[i + zeros - 1] = '0'
	return string(s)
}