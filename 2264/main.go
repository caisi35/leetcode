package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}

func largestGoodInteger(num string) string {
	for i := 9; i >= 0; i-- {
		s := strconv.Itoa(i)
		s3 := fmt.Sprintf("%s%s%s", s, s, s)
		if strings.Contains(num, s3) {
			return s3
		}
	}
	return ""
}