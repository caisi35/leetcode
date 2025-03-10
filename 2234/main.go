package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([]int{1,3,1,1}, 7, 6, 12, 1))	 // 14
	fmt.Println(maximumBeauty([]int{2,4,5,3}, 10, 5, 2, 6))	 // 30
}

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	for i := 0; i < n; i++ {
		if flowers[i] > target {
			flowers[i] = target
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(flowers)))
	var sum int64
	for _, flower := range flowers {
		sum += int64(flower)
	}
	var ans int64
	if int64(target)*int64(n) - sum <= newFlowers {
		ans = int64(full) * (int64(n))
	}
	pre, ptr := int64(0), 0
	for i := 0; i < n; i++ {
		if i != 0 {
			pre += int64(flowers[i-1])
		}
		if flowers[i] == target {
			continue
		}
		rest := newFlowers - (int64(target) * int64(i) - pre)
		if rest < 0 {
			break
		}
		for !(ptr >= i && int64(flowers[ptr]) * int64(n-ptr) - sum <= rest) {
			sum -= int64(flowers[ptr])
			ptr++
		}
		rest -= int64(flowers[ptr]) * int64(n- ptr) - sum
		ans = max(ans, int64(full) * int64(i) + int64(partial) * min(int64(flowers[ptr]) + rest / int64(n-ptr), int64(target) - 1)) 
	}
	return ans
}

func maximumBeauty2(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	n := len(flowers)
	sort.Ints(flowers)

	prefixSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + int64(flowers[i])
	}
	maxBeauty := int64(0)
	for k := 0; k <= n; k++ {
		required := int64(k) * int64(target) - (prefixSum[n] - prefixSum[n-k])
		if required > int64(newFlowers) {
			continue
		}
		remainint := int64(newFlowers) - required
		low := 0
		high := target - 1
		bestPartial := 0
		for low <= high {
			mid := (low + high) / 2
			idx := sort.Search(n-k, func(i int) bool {
				return flowers[i] > mid
			})
			meeded := int64(mid) *int64(idx) - (prefixSum[idx] - prefixSum[0])
			if meeded <= remainint {
				bestPartial = mid
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		totalBeauty := int64(k) * int64(full) + int64(bestPartial) * int64(partial)
		if totalBeauty > maxBeauty {
			maxBeauty = totalBeauty
		}
	}
	return maxBeauty
}