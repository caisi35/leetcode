package main

import (
	"fmt"
)

func main() {
	i := []int{5,6,7}
	foo(i)
	fmt.Println(i)
	fmt.Println(medianOfUniquenssArray([]int{1,2,3})) // 1
}

func foo(nums []int) {
	nums[0] = 18
}

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	median := (int64(n) * int64(n+1) / 2 + 1) / 2
	check := func(t int) bool {
		cnt := make(map[int]int)
		tot := int64(0)
		for i, j := 0, 0; i < n; i++ {
			cnt[nums[i]]++
			for len(cnt) > t {
				cnt[nums[j]]--
				if cnt[nums[j]] == 0 {
					delete(cnt, nums[j])
				}
				j++
			}
			tot += int64(i - j + 1)
		}
		return tot >= median
	}
	res := 0 
	lo, hi := 1, n
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(mid) {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}