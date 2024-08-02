package main

import "fmt"

func main() {
	fmt.Println(maxDivScore([]int{20,14,21,10}, []int{5,7,5}))
}

func maxDivScore(nums []int, divisors []int) int {
	ans := 0
	cnt := -1
	for _, divisor := range divisors {
		tmp := 0
		for _, num := range nums {
			if num % divisor == 0 {
				tmp++
			}
		}
		if tmp > cnt || (tmp == cnt && divisor < ans) {
			cnt = tmp
			ans = divisor
		}
	}
	return ans
}