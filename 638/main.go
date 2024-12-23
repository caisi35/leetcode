package main

import (
	"fmt"
)

func main() {
	fmt.Println(shoppingOffers([]int{2,3,4}, [][]int{{1,1,0,4}, {2,2,1,9}}, []int{1,2,1}))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	filterSpecial := [][]int{}
	for _, s := range special {
		totalCount, totalPrice := 0, 0
		for i, c := range s[:n] {
			totalCount += c
			totalPrice += c * price[i]
		}
		if totalCount > 0 && totalPrice > s[n] {
			filterSpecial = append(filterSpecial, s)
		}
	}
	dp := map[string]int{}
	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		if res, has := dp[string(curNeeds)]; has {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p
		}
		nextNeeds := make([]byte, n)
	outer:
		for _, s := range filterSpecial {
			for i, need := range curNeeds {
				if need < byte(s[i]) {
					continue outer
				}
				nextNeeds[i] = need - byte(s[i])
			}
			minPrice = min(minPrice, dfs(nextNeeds) + s[n])
		}
		dp[string(curNeeds)] = minPrice
		return
	}
	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}
	return dfs(curNeeds)
}