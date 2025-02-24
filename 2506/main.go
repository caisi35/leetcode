package main

import (
	"fmt"
)

func main() {
	fmt.Println(similarPairs([]string{"aabb", "ab", "ba"}))
	fmt.Println(similarPairs([]string{"nba", "cba", "dba"}))
}

func similarPairs(words []string) int {
	ans := 0
	n := len(words)
	m := []map[rune]int{}
	for i := 0; i < len(words); i++ {
		m = append(m, map[rune]int{})
	}
	for i, word := range words {
		for _, ch := range word {
			m[i][ch]++
		}
	}
	// fmt.Println(m)
	for i, word := range words {
		for j := i+1; j < n; j++ {
			if len(m[i]) != len(m[j]) {
				continue
			}
			f := true
			for _, ch := range word {
				if _, ok := m[j][ch]; !ok {
					f = false
					continue
				}
			}
			if f {
				ans++
			}
		}
	}

	return ans
}