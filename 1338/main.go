package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSetSize([]int{3,3,3,3,5,5,5,2,2,7}))
}

func minSetSize(arr []int) int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}
	occ := []int{}
	for _, v := range freq {
		occ = append(occ, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(occ)))
	cnt, ans := 0, 0
	for _, c := range occ {
		cnt += c
		ans++
		if cnt * 2 >= len(arr) {
			break
		}
	}
	return ans
}