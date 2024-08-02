package main

import (
	"fmt"
)

func change(s ...int) {
	fmt.Printf("%#v\n", s)
	s = append(s, 3)
	fmt.Printf("%#v\n", s)
}

func main() {
	slice := make([]int, 5) // 0 0 0 0 0
	fmt.Println(slice)
	slice[0] = 1          // 1 0 0 0 0
	slice[1] = 2          // 1 2 0 0 0
	change(slice...)      // 1 2 0 0 0 3	扩容，源slice没有变
	fmt.Println(slice)    // 1 2 0 0 0
	change(slice[0:2]...) //	1 2 3	未发生扩容，源slice改变
	fmt.Println(slice)    // 1 2 3 0 0
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
}

func maximumScore(nums []int, k int) int {
	left, right := k-1, k+1
	ans := 0
	for i := nums[k]; ; i-- {
		for left >= 0 && nums[left] >= i {
			left--
		}
		for right < len(nums) && nums[right] >= i {
			right++
		}
		ans = max(ans, (right-left-1)*i)
		if left == -1 && right == len(nums) {
			break
		}
	}
	return ans
}
