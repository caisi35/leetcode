package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(resultArray([]int{2, 1, 3, 3}))
	fmt.Println(resultArray([]int{5,14,3,1,2}))
	fmt.Println(resultArray([]int{3,3,3,3}))
	fmt.Println(resultArray2([]int{2, 1, 3, 3}))
	fmt.Println(resultArray2([]int{5,14,3,1,2}))
	fmt.Println(resultArray2([]int{3,3,3,3}))
}

func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if greaterCount(arr1, nums[i]) > greaterCount(arr2, nums[i]) {
			arr1 = append(arr1, nums[i])
		}
		if greaterCount(arr1, nums[i]) < greaterCount(arr2, nums[i]) {
			arr2 = append(arr2, nums[i])
		}
		if greaterCount(arr1, nums[i]) == greaterCount(arr2, nums[i]) {
			if len(arr1) == len(arr2) {
				arr1 = append(arr1, nums[i])
			} else if len(arr1) < len(arr2) {
				arr1 = append(arr1, nums[i])
			} else {
				arr2 = append(arr2, nums[i])
			}
		}
	}
	arr1 = append(arr1, arr2...)
	return arr1
}

func greaterCount(arr []int, val int) int {
	m := 0
	for _, i := range arr {
		if i > val {
			m++
		}
	}
	return m
}

type BinaryIndexedTree struct {
	tree []int
}

func NewBinaryIndexecTree(n int) *BinaryIndexedTree {
	return &BinaryIndexedTree{tree: make([]int, n+1)}
}

func (bit *BinaryIndexedTree) Add(i int) {
	for i < len(bit.tree) {
		bit.tree[i]++
		i += i & -i
	}
}

func (bit *BinaryIndexedTree) Get(i int) int {
	sum := 0
	for i > 0 {
		sum += bit.tree[i]
		i -= i & -i
	}
	return sum
}

func resultArray2(nums []int) []int {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	index := make(map[int]int)
	for i, num := range sortedNums {
		index[num] = i + 1
	}

	arr1, arr2 := []int{nums[0]}, []int{nums[1]}
	tree1, tree2 := NewBinaryIndexecTree(n), NewBinaryIndexecTree(n)
	tree1.Add(index[nums[0]])
	tree2.Add(index[nums[1]])
	for i := 2; i < n; i++ {
		count1 := len(arr1) - tree1.Get(index[nums[i]])
		count2 := len(arr2) - tree2.Get(index[nums[i]])
		if count1 > count2 || (count1 == count2 && len(arr1) <= len(arr2)) {
			arr1 = append(arr1, nums[i])
			tree1.Add(index[nums[i]])
		} else {
			arr2 = append(arr2, nums[i])
			tree2.Add(index[nums[i]])
		}
	}
	return append(arr1, arr2...)
}