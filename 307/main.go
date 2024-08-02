package main

import (
	"fmt"
)

func main() {
	c := Constructor([]int{1, 3, 5})
	fmt.Println(c.SumRange(0, 2))
	c.Update(1, 2)
	fmt.Println(c.SumRange(0, 2))
}

type NumArray struct {
	l []int
}


func Constructor(nums []int) NumArray {
	n := NumArray{l: nums}
	return n
}


func (this *NumArray) Update(index int, val int)  {
	this.l[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := 0
	for _, v := range this.l[left:right+1] {
		ans += v
	}
	return ans
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */