package main

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1,2}}))	// 2
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))  // -1
	fmt.Println(findJudge(3, [][]int{{1,3}, {2,3}}))	// 3
	fmt.Println(findJudge(3, [][]int{{1,3}, {2,3}, {3,1}}))	// -1
	fmt.Println(findJudge(4, [][]int{{1,3},{1,4},{2,3},{2,4},{4,3}}))	// 3
}

func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for _, t := range trust {
		in[t[1]]++
		out[t[0]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n - 1 && out[i] == 0 {
			return i
		}
	}
	return -1
}

func findJudge2(n int, trust [][]int) int {
	one := make(map[int]int)
	two := make(map[int]int)
	for _, t := range trust {
		one[t[0]] = t[0]
		two[t[1]] = t[1]
	}
	for k, v := range two {
		if _, ok := one[k]; ok {
			continue
		} else {
			return v
		}
	}
	return -1
}