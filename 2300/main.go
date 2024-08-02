package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	sort.Ints(potions)

	for i, v := range spells {
		fmt.Printf("(success-1)/v:%d   v:%d  (int(success) - 1) / v + 1:%d\n", (int(success) - 1) / v, v, (int(success) - 1) / v + 1)
		ans[i] = len(potions) - sort.SearchInts(potions, (int(success) - 1) / v + 1)
	}
	return ans
}
