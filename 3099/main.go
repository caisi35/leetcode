package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))	// 9
	fmt.Println(sumOfTheDigitsOfHarshadNumber(23))	// -1

}

func sumOfTheDigitsOfHarshadNumber(x int) int {
	ans := -1
	s := strconv.Itoa(x)
	t := 0
	for _, ch := range s {
		t += int(ch - '0')
		fmt.Println(t)
	}
	if x % t == 0 {
		ans = t
	}
	return ans
}