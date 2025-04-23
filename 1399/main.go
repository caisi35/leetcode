package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countLargestGroup(13))	// 4
}

func countLargestGroup(n int) int {
	ans := 0 
	m := map[int]int{}
	mx := 0
	for i:=1; i <= n; i++ {
		sn := strconv.Itoa(i)
		sum := 0
		for _, ch := range sn {
			sum += int(ch - '0')
		}
		m[sum]++
		mx = max(mx, m[sum])
	}
	// fmt.Println(m)
	for _, v := range m {
		if v == mx {
			ans++
		}
	}
	return ans
}