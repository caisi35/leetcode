package main

import (
	"fmt"
)

func main() {
	fmt.Println(validSubstringCount("bcca", "abc"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	count := make([]int, 26)
	for _, c := range word2 {
		count[c-'a']++
	}
	n := len(word1)
	preCount := make([][]int, n+1)
	for i := range preCount {
		preCount[i] = make([]int, 26)
	}
	for i := 1; i <= n; i++ {
		copy(preCount[i], preCount[i-1])
		preCount[i][word1[i-1]-'a']++
	}
	var res int64
	for l := 1; l <= n; l++ {
		r := get(l, n+1, preCount, count)
		res += int64(n-r+1)
	}
	return res
}

func get(l, r int, preCount [][]int, count []int) int {
	border := l
	for l < r {
		m := (l+r) >> 1
		f := true
		for i := 0; i < 26; i++ {
			if preCount[m][i]-preCount[border-1][i] < count[i] {
				f = false
				break
			}
		}
		if f {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}