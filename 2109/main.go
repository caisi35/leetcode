package main

import (
	"fmt"
)

func main() {
	fmt.Println(addSpaces("icodeinpython", []int{1,5,7,9}))
}

func addSpaces(s string, spaces []int) string {
	ans := []byte{}
	p := 0
	for i := 0; i < len(s); i++ {
		if p < len(spaces) && spaces[p] == i {
			p++
			ans = append(ans, ' ')
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

func addSpaces2(s string, spaces []int) string {
	ans := []rune{}
	m := map[int]struct{}{}
	for _, s := range spaces {
		m[s] = struct{}{}
	}
	for i, ch := range s {
		if _, ok := m[i]; ok {
			ans = append(ans,  ' ')
		}
		ans = append(ans, ch)
	}
	return string(ans)
}