package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getGoodIndices([][]int{
		{2,3,3,10},
		{3,3,3,1},
		{6,1,1,4},
	}, 2))
	fmt.Println(getGoodIndices([][]int{
		{45,34,41,32},
		{9,2,15,3},
		{31,12,21,24},
	}, 1))
	fmt.Println(getGoodIndices([][]int{
		{39,3,1000,1000},
	}, 17))
}

func getGoodIndices(veriables [][]int, target int) []int {
	ans := []int{}
	for i, v := range veriables {

		abp := math.Pow(float64(v[0]), float64(v[1]))
		ab := math.Mod(abp, float64(10))

		abc := math.Pow(float64(ab), float64(v[2]))
		abcd := math.Mod(abc, float64(v[3]))

		if float64(abcd) == float64(target) {
			ans = append(ans, i)
		}
	}

	// return ans

	// ans := []int{}
	for i, v := range veriables {
		if pow_mod(pow_mod(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		} 
	}
	return ans
}

func pow_mod(x, y, mod int) int {
	res := 1 
	for y > 0 {
		if (y & 1) == 1 {
			res = res * x % mod 
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}