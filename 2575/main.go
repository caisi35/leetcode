package main

import (
	"fmt"
	// "math"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	cur := 0
	for i, c := range word {
		cur = (cur * 10 + int(c - '0')) % m
		if cur == 0 {
			ans[i] = 1
		}
	}
	return ans
}

func main() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
	fmt.Println(divisibilityArray("998244353", 3))	// 1,1,0,0,0,1,1,0,0
	fmt.Println(divisibilityArray("1010", 10))	// 1,1,0,0,0,1,1,0,0
}
