package main

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1,1,2}))	// 5
	fmt.Println(numRabbits([]int{1,0,1,0,0}))	// 5
	fmt.Println(numRabbits([]int{0,0,1,1,1}))	// 6
}

func numRabbits(answers []int) int {
	m := map[int]int{}
	for _, a := range answers {
		m[a] += 1
	}
	ans := 0
	for k, v := range m {
		ans += (k+v) / (k+1) * (k+1)
	}
	return ans
}