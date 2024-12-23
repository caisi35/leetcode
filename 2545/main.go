package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortTheStudents([][]int{
		{3,4},
		{5,6},
	}, 0))
}

func sortTheStudents(score [][]int, k int) [][]int {
	s := [][]int{}
	for i := range score {
		s = append(s, []int{score[i][k], i})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] > s[j][0]
	})
	ans := [][]int{}
	for _, v := range s {
		ans = append(ans, score[v[1]])
	}
	return ans
}