package main

import "fmt"

func main() {
	fmt.Println(maximumSubsequenceCount("abcdefghijklmnoplrsuvwz", "ac"))
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	var res, cnt1, cnt2 int64
	for _, c := range text {
		if byte(c) == pattern[1] {
			res += cnt1
			cnt2++
		}
		if byte(c) == pattern[0] {
			cnt1++
		}
	}
	if cnt1 > cnt2 {
		return res + cnt1
	}
	return res + cnt2
}