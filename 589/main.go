package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
	preorder(&Node{})
	ans := []int{}
	t(&ans)
	fmt.Println(ans)
}

func t(ans *[]int) {
	*ans = append(*ans, 1)
	fmt.Println(ans)
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}
	g(root, &ans)
	return ans
}

func g(root *Node, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)

	for _, ch := range root.Children {
		g(ch, ans)
	}
}
