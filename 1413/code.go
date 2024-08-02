package main

func minStartValue(nums []int) int {
	ac, acm := 0, 0
	for _, num := range nums {
		ac += num
		acm = min(ac, acm)
	}
	return -acm+1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{-3,2,-3,4,2}
	minStartValue(nums)
}