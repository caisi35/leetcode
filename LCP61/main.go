package main 

import "fmt"

func main() {
	fmt.Println(temperatureTrend([]int{21,18,18,18,31}, []int{34,32,16,16,17}))
}

func temperatureTrend(temperatureA []int, temperatureB []int) int {
	ans := 0
	f := 0
	for i := 1; i < len(temperatureA); i++ {
		if c(temperatureA[i]-temperatureA[i-1], temperatureB[i]-temperatureB[i-1]) {
			f++
		} else {
			f = 0
		}
		ans = max(ans, f)
	}
	return ans
}

func c(a, b int) bool {
	if a == b {
		return a == b 
	} else if a > 0 && b > 0 {
		return true
	} else if a < 0 && b < 0 {
		return true
	}
	return false
}