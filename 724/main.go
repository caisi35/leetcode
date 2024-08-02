package main 

import "fmt"

func pivotIndex(nums []int) int {
	ans := 0 
	t := 0
	for _, v := range nums {
		t += v
	}

	for i, v := range nums {
		if 2 * ans + v == t {
			return i
		}
		ans += v
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
	fmt.Println(pivotIndex([]int{1,2,3}))
	fmt.Println(pivotIndex([]int{2,1,-1}))

}