package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(3))
}

func judgeSquareSum(c int) bool {
	for a:=0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}