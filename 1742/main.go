package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBalls(1, 10))
}

func countBalls(lowLimit int, highLimit int) int {
	ans := 0 
	m := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		s := strconv.Itoa(i)
		n := 0
		for _, v := range s {
			n += int(v - 48)
		}
		m[n]++
	}
	for _, v := range m {
		if v > ans {
			ans = v
		}
	}
	return ans
}