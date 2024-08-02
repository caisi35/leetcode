package main

import (
	"fmt"
)

func main() {
	type MyInt int
	var i int = 1
	var j MyInt = MyInt(i)
	fmt.Println(i, j)
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res int64
	var dfs func(int, int) int
	dfs = func(cur, fa int) int {
		peopleSum := 1
		for _, ne := range g[cur] {
			if ne != fa {
				peopleCnt := dfs(ne, cur)
				peopleSum, res = peopleSum+peopleCnt, res+int64((peopleCnt+seats-1) / seats)
			}
		}
		return peopleSum
	}
	dfs(0, -1)
	return res
}
