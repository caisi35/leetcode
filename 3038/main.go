package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{3, 2, 1, 4, 5}))
	fmt.Println(maxOperations([]int{3,2,6,1,4}))
	fmt.Println(maxOperations([]int{2,2,3,2,4,2,3,3,1,3}))
}

func maxOperations(nums []int) int {
    ans := 1
	s := nums[0] + nums[1]
    for i := 2; i < len(nums) - 1; i+=2 {
        if nums[i] + nums[i+1] == s {
			ans++
		} else {
			break
		}
    }
	return ans
}