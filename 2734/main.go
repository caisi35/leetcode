package main

import "fmt"

func main() {
	fmt.Println(smallestString("cbabc"))
}

func smallestString(s string) string {
	t := []byte(s)
	firstA := findFirstNotA(s)
	if firstA == len(s) {
		t[firstA-1] = 'z'
		return string(t)
	}

	secondA := findSecondA(s, firstA)
	res := []byte{}
	for i, ch := range t {
		if firstA <= i && i < secondA {
			res = append(res, ch - 1)
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func findSecondA(t string, firstA int) int {
	for i := firstA; i < len(t); i++ {
		if t[i] == 'a' {
			return i
		}
	}
	return len(t)
}

func findFirstNotA(t string) int {
	
	for i, ch := range t {
		if ch != 'a' {
			return i
		}
	}
	return len(t)
}