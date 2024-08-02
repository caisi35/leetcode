package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
	fmt.Println(count("1", "12", 1, 8))
	fmt.Println(count2("4179205230", "7748704426", 8, 46))

}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	ans := 0
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	for i := n1; i <= n2; i++ {
		// fmt.Println(i)
		s := 0
		for _, j := range strconv.Itoa(i) {
			s += int(j - 48)
		}
		if s >= min_sum && s <= max_sum {
			ans++
		}
	}
	return ans
}

const MOD = 1000000007
const N = 23
const M = 401

func count2(num1 string, num2 string, min_sum int, max_sum int) int {
	d := make([][]int, N)
	for i := range d {
		d[i] = make([]int, M)
		for j := range d[i] {
			d[i][j] = -1
		}
	}
	var dfs func(num string, i int, j int, limit bool) int 
	dfs = func(num string, i int, j int, limit bool) int {
		if j > max_sum {
			return 0
		}
		if i == -1 {
			if j >= min_sum {
				return 1
			}
			return 0
		}
		if !limit && d[i][j] != -1 {
			return d[i][j]
		}
		res := 0
		var up int
		if limit {
			up = int(num[i] - '0')
		} else {
			up = 9
		}
		for x := 0; x <= up; x++ {
			res = (res+dfs(num, i-1,j+x,limit&&x==up)) % MOD
		}
		if !limit {
			d[i][j] = res
		}
		return res
	}
	get := func(num string) int {
		num = reverse(num)
		return dfs(num, len(num)-1,0,true)
	}
	sub := func(num string) string {
		i := len(num) - 1
		arr := []byte(num)
		for arr[i] == '0' {
			i--
		}
		arr[i]--
		i++
		for ; i < len(num); i++ {
			arr[i] = '9'
		}
		return string(arr)
	}
	return (get(num2) - get(sub(num1)) + MOD) % MOD
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) -1; i < j; i,j=i+1,j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}