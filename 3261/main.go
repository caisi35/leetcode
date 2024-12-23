package main

import (
	"fmt"
)

func main() {
	fmt.Println(countKConstraintSubstrings("010101", 1, [][]int{
		{0,5},
		{1,4},
		{2,3},
	}))
}

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)
	count := [2]int{}
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}
	prefix := make([]int64, n+1)
	i := 0
	for j := 0; j < n; j++ {
		count[int(s[j]-'0')]++
		for count[0] > k && count[1] > k {
			count[int(s[i]-'0')]--
			right[i] = j
			i++
		}
		prefix[j+1] = prefix[j] + int64(j-i+1)
	}
	res := make([]int64, 0, len(queries))
	for _, query := range queries {
		l, r := query[0], query[1]
		i := min(right[l], r + 1)
		part1 := int64(i-l+1) * int64(i-l) / 2
		part2 := prefix[r+1] - prefix[i]
		res = append(res, part1 + part2)
	}
	return res
}