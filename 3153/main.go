package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sumDigitDifferences(nums []int) int64 {
	res := int64(0)
	for nums[0] > 0 {
		cnt := make([]int, 10)
		for i := 0; i < len(nums); i++ {
			cnt[nums[i] % 10]++
			nums[i] /= 10
		}
		for i := 0; i < 10; i++ {
			res += int64(len(nums) - cnt[i]) * int64(cnt[i])
		}
	}
	return res / 2
}

func main() {
	numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)
	fmt.Println(sumDigitDifferences([]int{13,23,12})) // 4
	foo()
}

func foo() {
	closed := make(chan struct{})
	go func() {
		time.Sleep(time.Second*6)
		close(closed)
		fmt.Println("goroutine end")
	}()
	// closed <- struct{}{}
	fmt.Printf("closed: %V\n", <-closed)

	time.Sleep(5 * time.Second)
	
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		fmt.Println(i)
		go func() {
			defer wg.Done()
			fmt.Printf("%V\n",<-start)
		}()
	}
	close(start)
	wg.Wait()

	select {
	case <-closed:
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}