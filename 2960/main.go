package main

import "fmt"

func main() {
	fmt.Println(countTestedDevices([]int{1, 1, 2, 1, 3}))
}

func countTestedDevices(batterPercentages []int) int {
	ans := 0
	for i, v := range batterPercentages {
		fmt.Println(i, v)
		if v > 0 && v - ans > 0 {
			ans++
		}
	}
	return ans
}