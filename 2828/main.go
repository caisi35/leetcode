package main 

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Println(isAcronym([]string{"alice","aob","charlie"}, "abc"))
}

func isAcronym(words []string, s string) bool {
	res := true
	if len(words) != len(s) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i][0] != s[i] {
			return false
		}
	}

	return res
}