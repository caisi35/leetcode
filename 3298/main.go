package main

import (
	"fmt"
)

func main() {
	fmt.Println(validSubstringCount("bcca", "abc"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	diff := make([]int, 26)
	for _, c := range word2 {
		diff[c-'a']--
	}
	cnt := 0
	for _, c := range diff {
		if c < 0 {
			cnt++
		}
	}
	var res int64
	l, r := 0, 0
	for l < len(word1) {
		for r < len(word1) && cnt > 0 {
			update(diff, int(word1[r]-'a'), 1, &cnt)
			r++
		}
		if cnt == 0 {
			res += int64(len(word1) - r + 1)
		}
		update(diff, int(word1[l]-'a'), -1, &cnt)
		l++
	}
	return res
}

func update(diff []int, c, add int, cnt *int) {
	diff[c] += add
	if add == 1 && diff[c] == 0 {
		*cnt--
	} else if add == -1 && diff[c] == -1 {
		*cnt++
	}
}