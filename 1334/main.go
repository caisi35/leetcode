package main

import (
	"fmt"
	"math"
)

func main() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
	fmt.Println(findTheCity(4, [][]int{
		{0, 1, 3},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}, 4))
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	ans := []int{math.MaxInt32 / 2, -1}
	mp := make([][]int, n)
	dis := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt32 / 2
			dis[i][j] = math.MaxInt32 / 2
			vis[i][j] = false
		}
	}
	for _, eg := range edges {
		from, to, weight := eg[0], eg[1], eg[2]
		mp[from][to], mp[to][from] = weight, weight
	}
	for i := 0; i < n; i++ {
		dis[i][i] = 0
		for j := 0; j < n; j++ {
			t := -1
			for k := 0; k < n; k++ {
				if !vis[i][k] && (t == -1 || dis[i][k] < dis[i][t]) {
					t = k
				}
			}
			for k := 0; k < n; k++ {
				dis[i][k] = min(dis[i][k], dis[i][t]+mp[t][k])
			}
			vis[i][t] = true
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ans[0] {
			ans[0] = cnt
			ans[1] = i
		}
	}
	return ans[1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
