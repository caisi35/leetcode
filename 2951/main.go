package main

import "fmt"

func main() {
	fmt.Println(findPeaks([]int{2, 4, 4}))
	fmt.Println(findPeaks([]int{1, 4, 3, 8, 5}))

}

func findPeaks(mountain []int) []int {
	ans := []int{}
	for i := 1; i < len(mountain) - 1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return ans
}