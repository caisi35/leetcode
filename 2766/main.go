package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relocateMarbles([]int{1,6,7,8}, []int{1,7,2}, []int{2,9,5}))
	fmt.Println(relocateMarbles([]int{3,4}, []int{4,3,1,2,2,3,2,4,1}, []int{3,1,2,2,3,2,4,1,1}))	// [1]

}

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] = 1
	}
	
	for i := range  moveFrom {
		if _, ok := m[moveFrom[i]]; ok {
			m[moveFrom[i]] = 0
			m[moveTo[i]] += 1
		}
	}
	ans := []int{}
	for k, v := range m {
		if v >= 1 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return ans
}