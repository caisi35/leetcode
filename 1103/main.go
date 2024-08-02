package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))	
	fmt.Println(distributeCandies(10, 3))	
	fmt.Println(distributeCandies(60, 4))	// [15,18,15,12]
}

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	t := 0
	for candies > 0 {
		for i := 1; i <= num_people; i++ {
			t++
			if candies <= t {
				ans[i-1] = ans[i-1] + candies
				candies = 0
				break
			} else {
				ans[i-1] = ans[i-1] + t
				candies -= t
			}

		}
	}

	return ans
}