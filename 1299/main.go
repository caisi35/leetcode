package main 

import (
	"fmt"
)

func main() {
	fmt.Println(replaceElements([]int{17,18,5,4,6,1}))
}

func replaceElements(arr []int) []int {
	fm := func (nums []int) int {
		mx := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > mx {
				mx = nums[i]
			}
		}
		return mx
	}
	m := fm(arr)
	arr[0] = m

	for i := 0; i < len(arr)-1; i++ {
		if m == arr[i] {
			m = fm(arr[i+1:])
		}
		arr[i] = m
	}

	arr[len(arr)-1] = -1
	return arr
}