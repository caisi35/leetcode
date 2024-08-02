package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time // 内嵌使用内嵌MarshalJSON方法
		N         int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s\n", m)
	fmt.Println(numberOfBoomerangs([][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}))
}

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		cnt := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[dis]++
		}
		for _, m := range cnt {
			ans += m * (m - 1)
		}
	}
	return ans
}
