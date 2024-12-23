package main

import "fmt"

func main() {
	fmt.Println(timeRequiredToBuy([]int{2,3,2}, 2))
	fmt.Println(timeRequiredToBuy([]int{5,1,1,1}, 0))
}

func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	for {
		for i, t := range tickets {
			if tickets[k] == 0 {
				return ans
			}
			if t > 0 {
				tickets[i]--
				ans++
			}
		}
	}
}