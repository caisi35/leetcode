package main

import "fmt"

func main() {
	fmt.Println(maximumLength("abcaba"))
}

func maximumLength(s string) int {
	n := len(s)
	cnt := make(map[byte][]int)
	for i, j := 0, 0; i < n; i = j {
		for j < n && s[j] == s[i] {
			j++
		}
		cnt[s[i]] = append(cnt[s[i]], j-i)
	}
	res := -1
	for _, vec := range cnt {
		lo, hi := 1, n - 2
		for lo <= hi {
			mid := (lo + hi) >> 1
			count := 0
			for _, x := range vec {
				if x >= mid {
					count += x - mid + 1
				}
			}
			if count >= 3 {
				if mid > res {
					res = mid
				}
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return res
}