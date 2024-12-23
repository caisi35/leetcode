package main

import (
	"fmt"
)

func main() {
	fmt.Println(winningPlayerCount(5, [][]int{{1,1},{2,4},{2,4},{2,4}}));	// 1
	fmt.Println(winningPlayerCount(5, [][]int{{1,1},{1,2},{1,3},{1,4}}));	// 0
	fmt.Println(winningPlayerCount(4, [][]int{{0,0},{1,0},{1,0},{2,1},{2,1},{2,0}})); //2
	fmt.Println(winningPlayerCount(2, [][]int{{0,8},{0,3}})); //1
	fmt.Println(winningPlayerCount(2, [][]int{{0,6},{0,6}})); //1
}

func winningPlayerCount(n int, pick [][]int) int {
	m := map[int][]int{}
	for _, i := range pick {
		if _, ok := m[i[0]]; ok {
			m[i[0]] = append(m[i[0]], i[1])
			continue
		}
		m[i[0]] = []int{i[1]}
	}
	ans := 0
	fmt.Println(m)
	for k, v := range m {
		mc := map[int]int{}
		for _, i := range v {
			if _, ok := mc[i]; ok {
				mc[i]++
				continue
			}
			mc[i] = 1
			
		}
		fmt.Println(mc)
		for _, v2 := range mc {
			if k + 1 <= v2 {
				ans++
				break
			}
		}
	}
	return ans
}