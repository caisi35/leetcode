package main

import "fmt"

func main() {
	fmt.Println(longestAwesome("3242415"))
}

func longestAwesome(s string) int {
	prefix := map[int]int{0: -1}
	ans := 0
	sequence := 0
	for j := 0; j < len(s); j++ {
		digit := int(s[j] - '0')
		sequence ^= (1 << digit)
		if preIndex, ok := prefix[sequence]; ok {
			ans = max(ans, j-preIndex)
		} else {
			prefix[sequence] = j
		}
		for k := 0; k < 10; k++ {
			if prevIndex, ok := prefix[sequence^(1<<k)]; ok {
				ans = max(ans, j-prevIndex)
			}
		}
	}
	return ans
}