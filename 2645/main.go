package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	fmt.Println(addMinimum("aaaabb"))

}

func addMinimum(word string) int {
	d := make([]int, len(word)+1)
	for i := 1; i <= len(word); i++ {
		d[i] = d[i-1] + 2
		if i > 1 && word[i-1] > word[i-2] {
			d[i] = d[i-1] - 1
		}
	}
	return d[len(word)]
}
