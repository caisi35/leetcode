package main

import (
	"fmt"
)

func main() {
	fmt.Println(countTexts("22233"))
}

func countTexts(pressedKeys string) int {
	m := 1000000007
	n := len(pressedKeys)
	dp3 := []int{1,1,2,4}
	dp4 := []int{1,1,2,4}
	for i := 4; i <= n; i++ {
		dp3 = append(dp3, (dp3[i-1] + dp3[i-2] + dp3[i-3]) % m)
		dp4 = append(dp4, (dp4[i-1] + dp4[i-2] + dp4[i-3] + dp4[i-4]) % m)
	}
	res := 1
	cnt := 1
	for i := 1; i < n; i++ {
		if pressedKeys[i] == pressedKeys[i-1] {
			cnt++
		} else {
			if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
				res = (res *dp4[cnt]) % m
			} else {
				res = (res * dp3[cnt]) % m
			}
			cnt = 1
		}
	}
	if pressedKeys[n-1] == '7' || pressedKeys[n-1] == '9' {
		res = (res * dp4[cnt]) % m
	} else {
		res = (res * dp3[cnt]) % m
	}
	return res
}