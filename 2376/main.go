package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSpecialNumbers(20))		// 19
	fmt.Println(countSpecialNumbers(135))		// 110
}

func countSpecialNumbers(n int) int {
	nStr := strconv.Itoa(n)
	memo := make(map[int]int)
	var dp func(int, bool, string) int
	dp = func(mask int, prefixSmaller bool, nStr string) int {
		if countOnes(mask) == len(nStr) {
			return 1
		}
		key := mask * 2
		if prefixSmaller {
			key++
		}
		if _, exists := memo[key]; !exists {
			res, lowerBound := 0, 0
			if mask == 0 {
				lowerBound = 1
			}
			upperBound := 9
			if !prefixSmaller {
				upperBound = int(nStr[countOnes(mask)] - '0')
			}
			for i := lowerBound; i <= upperBound; i++ {
				if (mask>>i) & 1 == 0 {
					res += dp(mask | (1 << i), prefixSmaller || i < upperBound, nStr)
				}
			}
			memo[key] = res
		}
		return memo[key]
	}
	res, prod := 0, 9
	for i := 0; i < len(nStr)-1; i++ {
		res += prod
		prod *= 9 - i
	}
	res += dp(0, false, nStr)
	return res
}

func countOnes(x int) int {
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}


func countSpecialNumbers2(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		f := true
		m := make(map[rune]bool)
		for _, ch := range s {
			if _, ok := m[ch]; ok {
				f = false 
				break
			} else {
				m[ch] = true
			}
		}
		if f {
			ans += 1
		}
	}
	return ans
}