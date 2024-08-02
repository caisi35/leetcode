package main

import (
	"fmt"
)

func main() {
	s1()
	s2()
	fmt.Println(finalString("abiString"))
}

func s1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) // 0 0 0 0 0 1 2 3
}

func s2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) // 1 2 3 4
}

func finalString(s string) string {
	ts := ""
	for _, v := range s {
		if v == 'i' {
			ts = reverse(ts)
		} else {
			ts = ts + string(v)
		}
	}
	return ts
}

func reverse(s string) string {
	t := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return string(t)
}
