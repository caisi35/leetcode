package main

import (
	"fmt"
)

func main() {
	fmt.Println(clearDigits("abc"))  // "abc"
	fmt.Println(clearDigits("ab34"))  // ""
	fmt.Println(clearDigits("d9"))  // ""
}

func clearDigits(s string) string {

	chs := []rune{}
	for _, ch := range s {
		if ch >= '0' && ch <= '9' && len(chs) > 0 {
			chs = chs[:len(chs)-1]
		} else {
			chs = append(chs, ch)
		}

	}
	// fmt.Println(chs)
	return string(chs)
}