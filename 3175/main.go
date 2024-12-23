package main

import "fmt"

func main() {
	fmt.Println(findWinningPlayer([]int{4,2,6,3,9}, 2))
}

func findWinningPlayer(skills []int, k int) int {
	i, lasti := 0, 0
	n := len(skills)
	cnt := 0
	for i < n {
		j := i + 1
		for j < n && skills[j] < skills[i] && cnt < k {
			cnt++
			j++
		}
		if cnt == k {
			return i
		}
		cnt = 1
		lasti = i
		i = j
	}
	return lasti
}