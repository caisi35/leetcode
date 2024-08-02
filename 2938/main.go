package main

import "fmt"

func main() {
	fmt.Println(minimumSteps("101"))
}

func minimumSteps(s string) int64 {
	ans := int64(0)
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			sum++
		} else {
			ans += int64(sum)
		}
	}
	return ans
}