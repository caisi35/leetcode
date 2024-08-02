package main

import (
	"math"
)

func main() {
	var a int8 = -1
	var b int8 = -128 / a

	println(b)

	// const a2 int8 = -1
	// var b2 int8 = -128 / a2

	// println(b2)
	println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
}

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		d[i] = math.MaxInt
	}
	mp := map[string]int{}
	for _, e := range dictionary {
		mp[e]++
	}
	for i := 1; i <= n; i++ {
		d[i] = d[i-1] + 1
		for j := i; j >= 0; j-- {
			if _, ok := mp[s[j:i]]; ok {
				d[i] = min(d[i], d[j])
			}
		}
	}
	return d[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
