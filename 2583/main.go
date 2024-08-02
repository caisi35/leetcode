package main

import (
	"fmt"
	"sort"
)

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age) // 28

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age) // 29
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age) // 28
	}()

	person = &Person{29}
	fmt.Println(kthLargestLevelSum(&TreeNode{Val: 2}, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	var levelSumArray []int64
	for len(q) > 0 {
		levelSum, size := int64(0), len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			levelSum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		levelSumArray = append(levelSumArray, levelSum)
	}
	if len(levelSumArray) < k {
		return -1
	}
	sort.Slice(levelSumArray, func(i, j int) bool {
		return levelSumArray[i] < levelSumArray[j]
	})
	return levelSumArray[len(levelSumArray)-k]
}

