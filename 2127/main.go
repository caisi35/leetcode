package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumInvitations([]int{2, 2, 1, 2}))
}

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indeg := make([]int, n)
	for _, x := range favorite {
		indeg[x]++
	}

	used := make([]bool, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		used[u] = true
		q = q[1:]
		v := favorite[u]
		f[v] = max(f[v], f[u]+1)
		indeg[v]--
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}
	ring, total := 0, 0
	for i := 0; i < n; i++ {
		if !used[i] {
			j := favorite[i]
			if favorite[j] == i {
				total += f[i] + f[j]
				used[i], used[j] = true, true
			} else {
				u, cnt := i, 0
				for true {
					cnt++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = max(ring, cnt)
			}
		}
	}
	return max(ring, total)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
