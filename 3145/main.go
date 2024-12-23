package main 

import (
	"fmt"
)

func main() {
	fmt.Println(findProductsOfElements([][]int64{{1,3,7}}))		// [4]
}

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, 0)
	for _, query := range queries {
		query[0]++
		query[1]++
		l := midCheck(query[0])
		r := midCheck(query[1])
		mod := int(query[2])

		res := int64(1)
		pre := countOne(l - 1)
		for j := 0; j < 60; j++ {
			if (1 << j) & l != 0 {
				pre++
				if pre >= query[0] && pre <= query[1] {
					res = res * (1 << j) % int64(mod)
				}
			}
		}

		if r > l {
			bac := countOne(r - 1)
			for j := 0; j < 60; j++ {
				if (1 << j) & r != 0 {
					bac++
					if bac >= query[0] && bac <= query[1] {
						res = res * (1 << j) % int64(mod)
					}
				}
			}
		}

		if r - l > 1 {
			xs := countPow(r - 1) - countPow(l)
			res = res * int64(powMod(2, xs, mod)) % int64(mod)
		}
		ans = append(ans, int(res))
	}
	return ans
}

func countOne(x int64) int64 {
	var res int64 = 0
	sum := 0
	for i := 60; i >= 0; i-- {
		if (1 << i) & x != 0 {
			res += int64(sum) * (1 << i)
			sum++
			if i > 0 {
				res += int64(i) * (1 << (i - 1))
			}
		}
	}
	res += int64(sum)
	return res
}

func countPow(x int64) int64 {
	var res int64 = 0
	sum := 0
	for i := 60; i >= 0; i-- {
		if (1 << i) & x != 0 {
			res += int64(sum) * (1 << i)
			sum += i
			if i > 0 {
				res += int64(i) * (int64(i) - 1) / 2 * (1 << (i - 1))
			}
		}
	}
	res += int64(sum)
	return res
}

func powMod(x int64, y int64, mod int) int {
	res := 1
	for y > 0 {
		if y & 1 != 0 {
			res = res * int(x) % mod
		}
		x = x * x % int64(mod)
		y >>= 1
	}
	return res
}

func midCheck(x int64) int64 {
	l, r := int64(1), int64(1e15)
	for l < r {
		mid := (l + r) >> 1
		if countOne(mid) >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}