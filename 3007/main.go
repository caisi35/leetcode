package main

import (
	"fmt"
	"math/bits"
	"time"
)

func main() {
	fmt.Println(findMaximumNumber(9, 1)) // 6
	c := make(chan struct{})
	go f(c)
	go f2(c)
	c <- struct{}{}
	time.Sleep(time.Second * 3)
}

func f(c chan struct{}) {
	fmt.Println("start")
	<- c
	fmt.Println("end")
}

func f2(c chan struct{}) {
	fmt.Println("start2")
	<- c
	fmt.Println("end2")
}

func findMaximumNumber(k int64, x int) int64 {
	l, r := int64(1), (k + 1) << x
	for l < r {
		m := (l + r + 1) / 2
		if accumulatePrice(x, m) > k {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func accumulateBitPrice(x int, num int64) int64 {
	period := int64(1) << x
	res := period / 2 * (num / period)
	if num % period >= period / 2 {
		res += num % period - (period / 2 - 1)
	}
	return res
}

func accumulatePrice(x int, num int64) int64 {
	res := int64(0)
	length := 64 - bits.LeadingZeros64(uint64(num))
	for i := x; i <= length; i += x {
		res += accumulateBitPrice(i, num)
	}
	return res
}