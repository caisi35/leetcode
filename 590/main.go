package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}


func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}


func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println(f1())   // 0
	fmt.Println(f2())	// 5
	fmt.Println(f3())	// 1
}

type Node struct {
    Val int
    Children []*Node
}

   
func postorder(root *Node) []int {
	ans := []int{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return 
		}
		for _, ch := range node.Children {
			dfs(ch)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}