package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg")) // 3
	loop()
}

func loop() {
	c := 0
Loop:
	for c < 10 {
		if c == 5 {
			continue Loop
		}
		c++
		print(c)
	}
	print(c)
}

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	d := make([]int, n+1)
	for i := range d {
		d[i] = 0x3f3f3f3f
	}
	d[0] = 0
	for i := 1; i <= n; i++ {
		maxCnt := 0
		occCnt := make(map[byte]int)
		for j := i; j >= 1; j-- {
			occCnt[s[j-1]]++
			if occCnt[s[j-1]] > maxCnt {
				maxCnt = occCnt[s[j-1]]
			}
			if maxCnt * len(occCnt) == (i-j+1) &&d[j-1] != 0x3f3f3f3f {
				if d[i] > d[j-1]+1 {
					d[i] = d[j-1]+1
				}
			}
		}
	}
	return d[n]
}