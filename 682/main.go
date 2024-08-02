package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))  // 30
	fmt.Println(calPoints([]string{"5","-2","4","C","D","9","+","+"}))   // 27

}

func calPoints(operations []string) int {
	ans := 0
	l := []int{}
	for _, s := range operations {
		if s == "D" {
			l = append(l, l[len(l)-1] * 2)
		} else if s == "C" {
			// f, _ := strconv.Atoi(ops[i-1])
			l = l[:len(l)-1]
		} else if s == "+" {
			l = append(l, l[len(l)-2] + l[len(l)-1])
		} else {
			v, _ := strconv.Atoi(s)
			l = append(l, v)
		}
	}
	for _, v := range l {
		ans += v
	}
	return ans
}