package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumDifference([]int{1,2,3,4,5}, 2))
}

func minimumDifference(nums []int, k int) int {
	n := len(nums)
	bits := make([]int, 31)
	for i := range bits {
		bits[i] = -1
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j <= 30; j++ {
			if nums[i] >> j & 1 == 1 {
				bits[j] = i
			}
		}
		pos := make([][2]int, 0)
		for j := 0; j <= 30; j++ {
			if bits[j] != -1 {
				pos = append(pos, [2]int{bits[j], j})
			}
		}
		sort.Slice(pos, func(i, j int) bool {
			return pos[i][0] > pos[j][0]
		})
		val := 0
		for j, p := 0, 0; j < len(pos); p = j {
			for j < len(pos) && pos[j][0] == pos[p][0] {
				val |= 1 << pos[j][1]
				j++
			}
			res = min(res, int(math.Abs(float64(val - k))))
		}
	}
	return res
}