package main     

import "fmt"

func main() {
	fmt.Println(maximumPrimeDifference([]int{4,2,9,5,3}))	// 3
	fmt.Println(maximumPrimeDifference([]int{4,8,2,8}))		// 0
	fmt.Println(maximumPrimeDifference([]int{2,2}))			// 1
	fmt.Println(maximumPrimeDifference([]int{1,7,7}))		// 1
	fmt.Println(maximumPrimeDifference([]int{2,2,10,6}))		// 1

}

func maximumPrimeDifference(nums []int) int {
	n := len(nums)
	fn := -1
	ln := -1
	for i := 0; i < n; i++ {
		if fn == -1 && prime(nums[i]) {
			fn = i
		}
		if prime(nums[i]) {
			ln = i
		}
	}
	return ln - fn
}

func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}