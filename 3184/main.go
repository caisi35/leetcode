package main

import "fmt"

func main() {
	fmt.Println(countCompleteDayPairs([]int{12,12,30,24,24}))
	fmt.Println(countCompleteDayPairs([]int{72, 48, 24, 3}))

}

func countCompleteDayPairs(hours []int) int {
	n := len(hours)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (hours[i] + hours[j]) % 24 == 0 {
				ans++
			}
		}
	}
	return ans
}