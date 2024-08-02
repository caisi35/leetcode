package main

import "fmt"

func main() {
	fmt.Println(allPossibleFBT(7))

	fmt.Println(allPossibleFBT(3))
}

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func allPossibleFBT(n int) []*TreeNode {
	var fullBianryTrees []*TreeNode
	if n % 2 == 0 {
		return fullBianryTrees
	}
	if n == 1 {
		fullBianryTrees = append(fullBianryTrees, &TreeNode{Val: 0})
		return fullBianryTrees
	}
	
	for i := 1; i < n; i += 2 {
		leftSubtrees := allPossibleFBT(i)
		rightSubtrees := allPossibleFBT(n - 1 - i)
		for _, leftSubtree := range leftSubtrees {
			for _, rightSubtree := range rightSubtrees {
				root := &TreeNode{Val: 0, Left: leftSubtree, Right: rightSubtree}
				fullBianryTrees = append(fullBianryTrees, root)
			}
		}
	}
	return fullBianryTrees
 }