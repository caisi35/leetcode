package main

import (
	"fmt"
	"net"
	"sort"
)

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func hIndex(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool {
		return citations[x] >= n-x
	})
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	fmt.Println(f1(), f2(), f3())
	fmt.Println(hIndex([]int{1, 3, 5, 9}))
	fmt.Println(net.ParseIP("127.0.0.1"))
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}
