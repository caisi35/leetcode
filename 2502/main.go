package main

import (
	"fmt"
)

func main() {
	c := Constructor(10)
	c.Allocate(1, 1)
	c.Allocate(1, 2)
	c.Allocate(1, 3)
	c.FreeMemory(2)
	c.Allocate(3, 4)
	c.Allocate(1, 1)
	fmt.Println(c.Allocate(1, 1))
}

type Allocator struct {
	n int
	memory []int
}

func Constructor(n int) Allocator {
	return Allocator {
		n: n,
		memory: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	count := 0
	for i := 0; i < this.n; i++ {
		if this.memory[i] != 0 {
			count = 0
		} else {
			count++
			if count == size {
				for j := i - count + 1; j <= i; j++ {
					this.memory[j] = mID
				}
				return i - count + 1
			}
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	count := 0
	for i := 0; i < this.n; i++ {
		if this.memory[i] == mID {
			count++
			this.memory[i] = 0
		}
	}
	return count
}