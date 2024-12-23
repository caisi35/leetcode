package main

import (
	"fmt"
	"sort"
)

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))	// 7
	fmt.Println(maxTotalReward([]int{1,1,3,3}))	// 4
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s)
	defer func() {
		fmt.Println(s[:1], len(s[:1]), cap(s[:1]))
		_ = append(s[:1], "x", "y")
	}()
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	m := rewardValues[len(rewardValues)-1]
	dp := make([]int, 2 * m)
	dp[0] = 1
	for _, x := range rewardValues {
		for k := 2 * x - 1; k >= x; k-- {
			if dp[k-x] == 1 {
				dp[k] = 1
			}
		}
	}
	x := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == 1 {
			x = i
		}
	}

	return x
}