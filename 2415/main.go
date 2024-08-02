package main

import (
	"fmt"
)

func main() {
	root := TreeNode{
		// [7,13,11]
		Val:   7,
		Left:  &TreeNode{Val: 13},
		Right: &TreeNode{Val: 11},
	}
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
	reverseOddLevels(&root)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	isOdd := 0
	fmt.Println(q)
	for len(q) > 0 {
		if isOdd != 0 {
			n := len(q)
			for i := 0; i < n/2; i++ {
				nodex, nodey := q[i], q[n-i-1]
				nodex.Val, nodey.Val = nodey.Val, nodex.Val
			}
		}
		tmp := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			if node.Left != nil {
				tmp = append(tmp, node.Left, node.Right)
			}
		}
		q = tmp
		isOdd ^= 1
	}
	return root
}
