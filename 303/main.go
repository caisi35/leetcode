package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
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

	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))

}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		ans += this.arr[i]
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
