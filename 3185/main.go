package main

import "fmt"

func main() {
	fmt.Println(countCompleteDayPairs([]int{12,12,23,43,24}))
}

func countCompleteDayPairs(hours []int) int64 {
	var ans int64
	m := make([]int, 24)
	for _, h := range hours {
		ans += int64(m[(24 - h % 24) % 24])
		m[h % 24]++
	}
	return ans
}