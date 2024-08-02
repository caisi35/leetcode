package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))

	var m = make(map[string]int)  //A
	m["a"] = 1
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}

}

func countSeniors(details []string) int {
	res := 0
	for _, d := range details {
		v, _ := strconv.Atoi(d[11:13])
		if v > 60 {
			res += 1
		}
	}

	return res
}
