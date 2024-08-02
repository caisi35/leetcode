package main

import (
	"fmt"

	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // 1 2 3 4 5
	fmt.Println("a = ", a) // 1 12 13 4 5

	obj := Constructor()
	obj.AddBack(2)
	param_1 := obj.PopSmallest()
	param_2 := obj.PopSmallest()
	param_3 := obj.PopSmallest()
	fmt.Println(param_1, param_2, param_3)
	obj.AddBack(1)
	param_4 := obj.PopSmallest()
	param_5 := obj.PopSmallest()
	param_6 := obj.PopSmallest()
	fmt.Println(param_4, param_5, param_6)
	fmt.Println(obj.m)
}

type SmallestInfiniteSet struct {
	m map[int]int
	// min int
	num int
	s treeset.Set
}


func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		m: make(map[int]int),
		// min: 0,
		num: 1,
	}
}

func get_m_min(m map[int]int) int {
	first := 0
	for k, _ := range m {
		first = k
		break
	}
	for k, _ := range m {
		if k != first && k < first {
			first = k
		}
	}
	return first
}


func (this *SmallestInfiniteSet) PopSmallest() int {
	ans := this.num

	if len(this.m) > 0 {
		min := get_m_min(this.m)
		if min < this.num {
			delete(this.m, min)
			ans = min
		} else {
			this.num += 1
		}
	} else {
		this.num += 1
	}
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int)  {
	if num < this.num {
		this.m[num] = 0
	}
}


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */