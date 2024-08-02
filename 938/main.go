package main

import "fmt"

func main() {
	fmt.Println(rangeSumBST(&TreeNode{Val: 10}, 7, 15))

	p := f()    //B
	fmt.Println(p.m) //print "foo"
}

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}  //A
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(node *TreeNode)
	ans := 0
	dfs = func(node *TreeNode) {
		if node.Val >= low && node.Val <= high {
			ans += node.Val
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}