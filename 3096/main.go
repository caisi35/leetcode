package main

import "fmt"

func main() {
	fmt.Println(minimumLevels([]int{1,0,1,0}))
}

func minimumLevels(possible []int) int {
	n := len(possible)
	tot := 0
	for _, v := range possible {
		tot += v
	}
	tot = tot * 2 - n
	pre := 0
	for i := 0; i < n - 1; i++ {
		if possible[i] == 1 {
			pre++
		} else {
			pre--
		}
		if 2 * pre > tot {
			return i + 1
		}
	}
	return -1
}