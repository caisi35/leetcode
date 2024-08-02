package main

import (
	"fmt"
	"time"
)

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func main() {
	// fmt.Println([...]int{1} == [2]int{1})
	// fmt.Println([]int{1} == []int{1})
	fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
}

func minDeletion(nums []int) int {
	ans := 0
	n := len(nums)
	check := true
	for i := 0; i < n-1; i++ {
		if check && nums[i] == nums[i+1] {
			ans++
		} else {
			check = !check
		}
	}
	if (n-ans)%2 != 0 {
		ans++
	}
	return ans
}
