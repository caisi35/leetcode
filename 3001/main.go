package main

import (
	"fmt"
)

func main() {
	fmt.Println(minMovesToCaptureTheQueen(1,1,8,8,2,3))
}

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e && (c != a || d <= min(b, f) || d >= max(b, f)) {
		return 1
	}
	if b == f && (d != b || c <= min(a, e) || c >= max(a, e)) {
		return 1
	}
	if abs(c-e) == abs(d-f) && ((c-e) * (b-f) != (a-e) * (d-f) || a < min(c, e) || a > max(c, e)) {
		return 1
	}
	return 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}