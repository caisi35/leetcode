package main

import "fmt"

func main() {
	fmt.Println(maxHeightOfTriangle(1, 1))	// 1
	fmt.Println(maxHeightOfTriangle(2, 4))	// 3
	fmt.Println(maxHeightOfTriangle(2, 1))	// 2
	fmt.Println(maxHeightOfTriangle(10, 1))	// 2

}

func maxHeightOfTriangle(red int, blue int) int {
	return max(f(red, blue), f(blue, red))
}

func f(red, blue int) int {
	for i := 1; ; i++ {
		if i % 2 == 1 {
			red -= i
			if red < 0 {
				return i - 1
			}
		} else {
			blue -= i
			if blue < 0 {
				return i - 1
			}
		}
	}
}