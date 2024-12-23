package main

import (
	"fmt"
)

func main() {
	fmt.Println(stableMountains([]int{10,1,10,1,10}, 10))
	fmt.Println(stableMountains([]int{10,1,10,1,10}, 3))
}

func stableMountains(height []int, threshold int) []int {
	ans := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			ans = append(ans, i)
		} 
	}
	return ans
}