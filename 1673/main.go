package main

import "fmt"

func main() {
	fmt.Println(mostCompetitive([]int{3,5,2,6}, 2))
}

func mostCompetitive(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	for i, x := range nums {
		for len(res) > 0 && len(nums) - i + len(res) > k && res[len(res)-1] > x {
			res = res[:len(res)-1]
		}
		res = append(res, x)
	}
	return res[:k]
}