package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(sumOfPowers([]int{1,2,3,4}, 3))
}

func sumOfPowers(nums []int, k int) int {
	const mod = 1e9 + 7
	const inf = 0x3f3f3f3f
	n := len(nums)
	res := 0
	d := make([][]map[int]int, n)
	for i := range d {
		d[i] = make([]map[int]int, k + 1)
		for j := range d[i] {
			d[i][j] = make(map[int]int)
		}
	}
	sort.Ints(nums)
	for i := 0 ; i < n; i++ {
		d[i][1][inf] = 1
		for j := 0; j < i; j ++ {
			diff := int(math.Abs(float64(nums[i] - nums[j])))
			for p := 2; p <= k; p++ {
				for v, cnt := range d[j][p-1] {
					d[i][p][min(diff, v)] = (d[i][p][min(diff, v)] + cnt) % mod
				}
			}
		}
		for v, cnt := range d[i][k] {
			res = (res + v * cnt % mod) % mod
		}
	}
	return res
}