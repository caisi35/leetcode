package main

import (
	"math"
	"fmt"
)

func maxStrength(nums []int) int64 {
	n, z, p := 0, 0, 0
	prod := 1
	mn := math.MinInt32
	for _, num := range nums {
		if num < 0 {
			n++
			prod *= num
			if num > mn {
				mn = num
			}
		} else if num == 0 {
			z++
		} else {
			prod *= num
			p++
		}
	}
	if n == 1 && z == 0 && p == 0 {
		return int64(nums[0])
	}
	if n <= 1 && p == 0 {
		return int64(0)
	}
	if prod < 0 {
		return int64(prod / mn)
	} else {
		return int64(prod)
	}
}

func main() {
	fmt.Println(maxStrength([]int{-4, -5, -4})) // 20
}