package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTaskAssign([]int{12,32,3,4}, []int{12,3,5,3}, 3, 5))
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	check := func(mid int) bool {
		p := pills
		ws := list.New()
		ptr := m - 1
		for i := mid - 1; i >= 0; i-- {
			for ptr >= m - mid && workers[ptr]+strength >= tasks[i] {
				ws.PushFront(workers[ptr])
				ptr--
			}
			if ws.Len() == 0 {
				return false
			}
			if ws.Back().Value.(int) >= tasks[i] {
				ws.Remove(ws.Back())
			} else {
				if p == 0 {
					return false
				}
				p--
				ws.Remove(ws.Front())
			}
		}
		return true
	}
	left, right, ans := 1, min(n, m), 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}