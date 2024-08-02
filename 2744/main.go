package main

import "fmt"

func hello(num ...int) {  
    num[0] = 18
}

func main() {  
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
	fmt.Println(maximumNumberOfStringPairs([]string{"ku","dd","gu","uk"}))
}

func maximumNumberOfStringPairs(words []string) int {
	ans := 0
	n := len(words)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if words[i] == reserve(words[j]) {
				ans++
			}
		}
	}
	return ans
}

func reserve(s string) string {
	a := []byte{}
	for i := len(s)-1; i >= 0; i-- {
		a = append(a, s[i])
	}
	return string(a)
}