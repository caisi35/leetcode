package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(5, 2))
}

func distributeCandies(n int, limit int) int {
	ans := 0;
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i + j > n {
				break
			}
			if n - j - i <= limit {
				ans += 1;
			}
		}
	}
	return ans
}