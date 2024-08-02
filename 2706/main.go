package main

import (
	"fmt"
	"sort"
)

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
	fmt.Println(buyChoco([]int{1, 2, 2}, 3))
}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	fmt.Println(prices)
	sum := money - (prices[0] + prices[1])
	if sum  > 0 {
		return sum
	} else if sum == 0 {
		return 0
	}
	return money
}
