package main

import (
	"fmt"
	"math/bits"
	"time"
)

func main() {
	fmt.Println(minEnd(3, 4)) // 6

	done := make(chan error, 1)
	go func() {
		time.Sleep(time.Second * 6)
		done <- fmt.Errorf("go err")
	}()
	select {
	case err := <-done:
		fmt.Println("error", err)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
	fmt.Println("end")
}

func minEnd(n int, x int) int64 {
	bitCount := 128 - bits.LeadingZeros(uint(n)) - bits.LeadingZeros(uint(x))
	res := int64(x)
	m := int64(n) - 1
	j := 0
	for i := 0; i < bitCount; i++ {
		if res&(1<<i) == 0 {
			if m&(1<<j) != 0 {
				res |= 1 << i
			}
			j++
		}
	}
	return res
}
