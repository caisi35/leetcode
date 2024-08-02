package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxmiumScore([]int{1,2,8,9}, 3))	// 18
	fmt.Println(maxmiumScore([]int{3,3,1}, 1))	// 0
	fmt.Println(maxmiumScore([]int{1, 3, 4, 5}, 4)) // 0
	fmt.Println(maxmiumScore([]int{9,5,9,1,6,10,3,4,5,1}, 2)) // 18
}

func maxmiumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	ans := 0
	odd := -1
	even := -1
	for i := 0; i < cnt; i++ {
		ans += cards[i]
		if cards[i] % 2 == 0 {
			even = cards[i]
		} else {
			odd = cards[i]
		}
	}

	if ans%2 == 0 {
		return ans
	}
	tmp := 0
	for i := cnt; i < len(cards); i++ {
		if cards[i] % 2 == 1 {
			if even != -1 {
				tmp = max(tmp, ans - even + cards[i])
			}
		} else {
			if odd != -1 {
				tmp = max(tmp, ans - odd + cards[i])
			}
		}
	}
	return tmp
}
