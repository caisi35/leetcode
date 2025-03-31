package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCycle([]int{3,3,4,2,3})) // 3
}

func longestCycle(edges []int) int {
	n := len(edges)
	label := make([]int, n)
	current_label, ans := 0, -1
	for i := 0; i < n; i++ {
		if label[i] != 0 {
			continue
		}
		pos, start_label := i, current_label
		for pos != -1 {
			current_label++
			if label[pos] != 0 {
				if label[pos] > start_label {
					if current_label - label[pos] > ans {
						ans = current_label - label[pos]
					}
				}
				break
			}
			label[pos] = current_label
			pos = edges[pos]
		}
	}
	return ans
}