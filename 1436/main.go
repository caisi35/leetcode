package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}))
}

func destCity(paths [][]string) string {
	ans := ""
	s := map[string]bool{}
	for _, i := range paths {
		s[i[0]] = true
	}

	for _, i := range paths {
		if _, ok := s[i[1]]; !ok {
			ans = i[1]
			break
		}
	}
	return ans
}