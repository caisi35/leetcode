package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	root := &TreeNode{Val: -1, Left: nil, Right: &TreeNode{Val: -1, Left: nil, Right: nil}}
	obj := Constructor(root)
	fmt.Printf("%#v\n%#v\n", obj.node, obj.node.Right)
	fmt.Println(obj.Find(1))
	fmt.Println(obj.Find(2))

	ch3 := make(chan int, 10) 
	go func ()  {
		var i = 1
		for {
			i++
			ch3 <- 1
		}
	}()

	for {
		select {
		// case x := <- ch3:
			// time.Sleep(time.Second)
			// fmt.Println(x)
		case <- time.After(1 * time.Second):
			println(time.Now().Unix())
		}
	}

	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}

	var wg sync.WaitGroup

	ch2 := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case ch2 <- getVal(1):
			fmt.Println("in first case")
		case ch2 <- getVal(2):
			fmt.Println("in second case")
		default:
			fmt.Println("default")
		}
	}()
	fmt.Println("The val:", <- ch2)

	foo := make(chan int) 
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <- bar:
		default:
			fmt.Println("default")
		}
	}()
	wg.Wait()
}

func getVal(i int) int {
	fmt.Println("getVal, i=", i)
	return i
}

func talk(msg string, sleen int) <- chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleen) * time.Microsecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <- chan string) <- chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <- input1:
			case ch <- <- input2:
			// case t := <- input1:
			// 	ch <- t
			// case t := <- input2:
			// 	ch <- t
			}
		}
	}()
	return ch
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	node *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	var dfs func(node *TreeNode, x int)
	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		node.Val = x
		dfs(node.Left, x * 2 + 1)
		dfs(node.Right, x * 2 + 2)		
	}
	dfs(root, 0)
	return FindElements{node: root}
}

func (this *FindElements) Find(target int) bool {
	ans := false
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return 
		}
		if node.Val == target {
			ans = true
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(this.node)
	return ans
}
