package main

import (
	"fmt"
	"math"
	"strconv"
)

func categorizeBox(length int, width int, height int, mass int) string {
	ans := make(map[string]string)
	tj := length * width * height >= int(math.Pow10(9))
	wd := int(math.Pow10(4))
	if length >= wd || width >= wd || height >= wd || tj {
		ans["Bulky"] = ""
	}
	if mass >= 100 {
		ans["Heavy"] = ""
	}
	_, ok1 := ans["Bulky"]
	_, ok2 := ans["Heavy"]
	if ok1 && ok2 {
		return "Both"
	}
	if !ok1 && !ok2 {
		return "Neither"
	}
	if ok1 && !ok2 {
		return "Bulky"
	}
	if ok2 && !ok1 {
		return "Heavy"
	}
	
	return ""
}

func main() {
	i := 65
	n, err := strconv.Atoi("65")
	fmt.Println(string(i), n, err)
	fmt.Println(categorizeBox(200, 50, 800, 50))

	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(s, cap(s))
	fmt.Println(a, cap(a))
	fmt.Println(b, cap(b))
	fmt.Println(c, cap(c))

}