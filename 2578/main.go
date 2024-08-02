package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}

func splitNum(num int) int {
	str_num := []byte(strconv.Itoa(num))

	sort.Slice(str_num, func(i, j int) bool {
		return str_num[i] < str_num[j]
	})

	n1, n2 := 0, 0
	for i := 0; i < len(str_num); i++ {
		if i%2 == 0 {
			n1 = n1*10 + int(str_num[i]-'0')
		} else {
			n2 = n2*10 + int(str_num[i]-'0')
		}
		fmt.Println(n1, " - ", n2)
	}
	return n1 + n2
}
