package main

import (
	"fmt"
)

var a bool = true
func main() {
	fmt.Println(maximumMinutes([][]int{
		{0,2,0,0,0,0,0},
		{0,0,0,2,2,1,0},
		{0,2,0,0,1,2,0},
		{0,0,2,2,2,0,2},
		{0,0,0,0,0,0,0},
	}))
	
	defer func(){
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func(){
		fmt.Println("3")
	}()
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1},{0, 1}}
var INF int = 1e9

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fireTime := make([][]int, m)
	for i := 0; i < m; i++ {
		fireTime[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fireTime[i][j] = INF
		}
	}
	bfs := func() {
		q := [][]int{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					q = append(q, []int{i, j})
					fireTime[i][j] = 0
				}
			}
		}
		time := 1
		for len(q) > 0 {
			tmp := q
			q = [][]int{}
			for _, p := range tmp {
				cx, cy := p[0], p[1]
				for _, d := range dirs {
					nx, ny := cx + d[0], cy + d[1]
					if nx >= 0 && ny >= 0 && nx < m && ny < n {
						if grid[nx][ny] == 2 || fireTime[nx][ny] != INF {
							continue
						}
						q = append(q, []int{nx, ny})
						fireTime[nx][ny] = time
					}
				}
			}
			time += 1
		}
	}

	bfs()
	check := func(stayTime int) bool {
		visit := make([][]bool, m)
		for i := 0; i < m; i++ {
			visit[i] = make([]bool, n)
		}
		q := [][]int{}
		q = append(q, []int{0, 0, stayTime})
		for len(q) > 0 {
			tmp := q
			q = [][]int{}
			for _, p := range tmp {
				cx, cy, time := p[0], p[1], p[2]
				for _, d := range dirs {
					nx, ny := cx + d[0], cy + d[1]
					if nx >= 0 && ny >= 0 && nx < m && ny <n {
						if visit[nx][ny] || grid[nx][ny] == 2{
							continue
						}
						if nx == m - 1 && ny == n - 1 {
							return fireTime[nx][ny] >= time + 1
						}
						if fireTime[nx][ny] > time + 1 {
							q = append(q, []int{nx, ny, time + 1})
							visit[nx][ny] = true
						}
					}
				}
			}
		}
		return false
	}
	ans := -1
	low, high := 0, m * n
	for low <= high {
		mid := low + (high - low) / 2
		if check(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if ans >= m * n {
		return INF
	}
	return ans
}