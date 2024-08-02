package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p :=1
	incr(&p)
	fmt.Println(p)
	fmt.Println(canWinNim(4))

}

func canWinNim(n int) bool {
	return n % 4 != 0
}