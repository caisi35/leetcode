package main

import (
	"fmt"
)

func main() {
	h := ListNode{}
	r := TreeNode{}
	fmt.Println(isSubPath(&h, &r))
}

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(rt *TreeNode, head *ListNode) bool {
	if head == nil {
		return true
	}

	if rt == nil {
		return false
	}

	if rt.Val != head.Val {
		return false
	}
	return dfs(rt.Left, head.Next) || dfs(rt.Right, head.Next)
}
