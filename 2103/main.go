package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
}

func countPoints(rings string) int {
	m := make(map[int][]string, 10)
	for i := 0; i < len(rings); i++ {
		k, _ := strconv.Atoi(string(rings[i+1]))
		col := string(rings[i])
		c_m := make(map[string]int, 3)
		for _, v := range m[k] {
			c_m[v] = 0
		}
		if _, ok := c_m[col]; !ok {
			m[k] = append(m[k], col)
		}
		
		i = i + 1
	}
	// fmt.Println(m)
	ans := 0
	for _, v := range m {
		if len(v) == 3 {
			ans++
		}
	}
	return ans
}