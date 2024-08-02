package main

import (
	"fmt"
	"math"
)

func minimumMoves(nums []int, k int, maxChanges int) int64 {
	n := len(nums)
	f := func(i int) int {
		x := nums[i]
		if i - 1 >= 0 {
			x += nums[i - 1]
		}
		if i + 1 < n {
			x += nums[i + 1]
		}
		return x
	}

	indexSum, sum := make([]int64, n + 1), make([]int64, n + 1)
	for i := 0; i < n; i++ {
		indexSum[i + 1] = indexSum[i] + int64(nums[i]) * int64(i)
		sum[i + 1] = sum[i] + int64(nums[i])
	}
	var res int64 = math.MaxInt64
	for i := 0; i < n; i++ {
		if f(i) + maxChanges >= k {
			if k <= f(i) {
				res = min(res, int64(k - nums[i]))
			} else {
				res = min(res, int64(2 * k - f(i) - nums[i]))
			}
			continue
		}
		left, right := 0, n
		for left <= right {
			mid := (left + right) / 2
			i1, i2 := max(i - mid, 0), min(i + mid, n - 1)
			if sum[i2 + 1] - sum[i1] >= int64(k - maxChanges) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		i1, i2 := max(i - left, 0), min(i + left, n - 1)
		if sum[i2 + 1] - sum[i1] > int64(k - maxChanges) {
			i1++
		}
		count1, count2 := sum[i + 1] - sum[i1], sum[i2 + 1] - sum[i + 1]
		res = min(res, indexSum[i2 + 1] - indexSum[i + 1] - int64(i) * count2 + int64(i) * count1 - (indexSum[i + 1] - indexSum[i1]) + 2 * int64(maxChanges))
	}
	return res
}

func main() {
	fmt.Println(minimumMoves([]int{0,0,0,0}, 2, 3))
}