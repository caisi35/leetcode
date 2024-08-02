package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s\n", m)
	fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
}

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		f, t := map[int]int{}, map[int]int{}
		for fi := 0; fi <= i; fi++ {
			if _, ok := f[nums[fi]]; !ok {
				f[nums[fi]] = nums[fi]
			}
		}
		for ti := i + 1; ti < n; ti++ {
			if _, ok := t[nums[ti]]; !ok {
				t[nums[ti]] = nums[ti]
			}
		}
		ans[i] = len(f) - len(t)
	}
	return ans
}
