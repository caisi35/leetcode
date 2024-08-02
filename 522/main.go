package main

import "fmt"

func main() {
	fmt.Println(findLUSlength([]string{
		"aba", "cdc", "eae",
	}))
}

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		for j, t := range strs {
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if ans < len(s) {
			ans = len(s)
		}
	}
	return ans
}

func isSubseq(s, t string) bool {
	pts := 0
	for ptt := range t {
		if s[pts] == t[ptt] {
			if pts++; pts == len(s) {
				return true
			}
		}
	}
	return false
}
