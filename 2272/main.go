package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestVariance("abcasdfea"))
}

func largestVariance(s string) int {
	pos := make(map[rune][]int)
	for i, ch := range s {
		pos[ch] = append(pos[ch], i)
	}
	ans := 0
	for c0, pos0 := range pos {
		for c1, pos1 := range pos {
			if c0 != c1 {
				i, j := 0, 0
				f, g := 0, -1<<63
				for i < len(pos0) || j < len(pos1) {
					if j == len(pos1) || (i < len(pos0) && pos0[i] < pos1[j]) {
						f, g = max(f, 0)+1, g+1
						i++
					} else {
						f, g = max(f, 0)-1, max(f, g, 0)-1
						j++
					}
					ans = max(ans, g)
				}
			}
		}
	}
	return ans
}

func largestVariance2(s string) int {
	n := len(s)
	maxDiff := 0
	for i := 0; i < n; i++ {
		freq := make(map[byte]int)
		for j := 0; j < n; j++ {
			freq[s[j]]++
			maxCount, minCount := 0, n
			for _, count := range freq {
				if count > maxCount {
					maxCount = count
				}
				if count < minCount {
					minCount = count
				}
			}
			diff := maxCount - minCount
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	return maxDiff
}
