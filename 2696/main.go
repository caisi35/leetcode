package main

import "fmt"

func main() {
	fmt.Println(t2("ABFCACDB"))
}

func minLength(s string) int {
	ans := len(s)
	for {
		for i := 0; i < len(s)-1; i++ {
			// fmt.Println(s[i]=='A', s[i+1] == 'B')
			if (s[i] == 'A' && s[i+1] == 'B') || (s[i] == 'C' && s[i+1] == 'D') {
				s = s[:i] + s[i+2:]
			}
		}
		fmt.Println(ans)
		if len(s) == ans {
			break
		}
		ans = len(s)
	}
	return ans
}

func t2(s string) int {
	var stack []byte
	for i := range s {
		stack = append(stack, s[i])
		m := len(stack)
		if m >= 2 && (string(stack[m-2:]) == "AB" || string(stack[m-2:]) == "CD") {
			stack = stack[:m-2]
		}
	}
	return len(stack)
}
