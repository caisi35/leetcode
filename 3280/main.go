package main

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	ans := []string{}
	for _, v := range strings.Split(date, "-") {
		i, _ := strconv.Atoi(v)
		ans = append(ans, strconv.FormatInt(int64(i), 2))
	}
	return strings.Join(ans, "-")
}

func main() {
	println(convertDateToBinary("2080-02-29"))
}