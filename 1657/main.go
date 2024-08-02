package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2]) // A B C
	fmt.Println(s2[0], s2[1], s2[2]) // C C C
	fmt.Println(closeStrings("abc", "bca"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	count1, count2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		count1[c-'a']++
	}
	for _, c := range word2 {
		count2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if count1[i] > 0 && count2[i] == 0 || count1[i] == 0 && count2[i] > 0 {
			return false
		}
	}
	sort.Ints(count1)
	sort.Ints(count2)
	return reflect.DeepEqual(count1, count2)
}
