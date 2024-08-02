package main

import "fmt"

func main() {
	fmt.Println(minimumOperations("2245047"))	// 2
}

func minimumOperations(num string) int {
	n := len(num)
	f1 := false
	z := false
	for i := n-1; i >= 0; i-- {
		t := num[i] - 48
		if t == 0 || t == 5 {
			if z {
				return n - i - 2
			} 
			if t == 5 {
				f1 = true
			} else if t == 0 {
				z = true
			}
		}

		if t == 2 || t == 7 {
			if f1 {
				return n - i -2
			}
		}
	}
	if z {
		return n-1
	}
	return n
}