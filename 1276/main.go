package main

import "fmt"

const cl = 100

var bl = 123

func main() {
	println(&bl, bl)
	// println(&cl,cl)
	fmt.Println(numOfBurgers(16, 7))
}

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || tomatoSlices < cheeseSlices*2 || cheeseSlices*4 < tomatoSlices {
		return nil
	}
	return []int{tomatoSlices/2 - cheeseSlices, cheeseSlices*2 - tomatoSlices/2}
}