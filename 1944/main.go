package main

import "fmt"

func main() {
	// var x = nil
	var x interface{} = nil
	// var x string = nil
	// var x error = nil

	fmt.Println(x)
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
	fmt.Println(can2([]int{10, 6, 8, 5, 11, 9}))
}

func canSeePersonsCount(heights []int) []int {
	ans := make([]int, len(heights))
	for i := 0; i < len(heights)-1; i++ {
		m := 0
		for j := i + 1; j < len(heights); j++ {

			if heights[j] > heights[i] {
				ans[i]++
				break
			} else if heights[j] > m {
				ans[i]++
			}
			m = max(m, heights[j])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func can2(heights []int) []int {
	n := len(heights)
	stack := make([]int, 0)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		h := heights[i]
		for len(stack) > 0 && stack[len(stack)-1] <= h {
			stack = stack[:len(stack)-1]
			res[i] += 1;
		}
		if len(stack) > 0 {
			res[i] += 1
		}
		stack = append(stack, h)
	}
	return res
}
