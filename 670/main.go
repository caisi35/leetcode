package main

import (
	"fmt"
	"strconv"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func main() {
	var w worker = person{}
	fmt.Println(w)
	fmt.Println(maximumSwap(2736))
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	m := num
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			n, _ := strconv.Atoi(string(s))
			if n > m {
				m = n
			}
			s[i], s[j] = s[j], s[i]

		}
	}
	return m
}