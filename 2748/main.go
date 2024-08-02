package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBeautifulPairs([]int{2, 5, 1, 4}))
}

func countBeautifulPairs(nums []int) int {
	strs := []string{}
	ans := 0
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	for i := range len(strs) {
		for j := i+1; j < len(strs); j++ {
			first := strs[i][0] - '0'
			last := strs[j][len(strs[j])-1] - '0'
			if gcd(int(first), int(last) % 10) {
				ans++
			}
		}
	}

	return ans
}

func gcd(a, b int) bool {
	for b != 0 {
		a, b = b, a % b
	}
	return a == 1
}