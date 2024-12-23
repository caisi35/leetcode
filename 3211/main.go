package main

import (
	"fmt"
)

func main() {
	fmt.Println(validStrings(3))
}

func validStrings(n int) []string {
	res := []string{}
	var dfs func(arr []rune)
	dfs = func(arr []rune) {
		if len(arr) == n {
			res = append(res, string(arr))
			return
		}
		if len(arr) == 0 || arr[len(arr)-1] == '1' {
			arr = append(arr, '0')
			dfs(arr)
			arr = arr[:len(arr)-1]
		}
		arr = append(arr, '1')
		dfs(arr)
		arr = arr[:len(arr)-1]
	}
	dfs([]rune{})
	return res
}