package main

import "fmt"

func main() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)) // 4
	fmt.Println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))    // 0
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					// fmt.Println(arr[i], arr[j], arr[k])
					ans++
				}
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
