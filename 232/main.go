package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	// 10 1 2 3     传 "1" 1 3
	a = 0
	defer calc("2", a, calc("20", a, b))
	// "20" 0 2 2	传 "2" 0 2
	b = 1
	// "2" 0 2 2
	// "1" 1 3 4

	x := 1
	obj := Constructor()
	obj.Push(x)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.Pop()) // 1
	obj.Push(5)
	fmt.Println(obj.Pop()) // 2
	fmt.Println(obj.Pop()) // 3
	fmt.Println(obj.Pop()) // 4
	fmt.Println(obj.Pop()) // 5

	fmt.Println(obj.Empty()) //true

	obj2 := Constructor2()
	obj2.Push(x)
	obj2.Push(2)
	obj2.Push(3)
	obj2.Push(4)
	fmt.Println(obj2.Pop()) // 1
	obj2.Push(5)
	fmt.Println(obj2.Pop()) // 2
	fmt.Println(obj2.Pop()) // 3
	fmt.Println(obj2.Pop()) // 4
	fmt.Println(obj2.Pop()) // 5

	fmt.Println(obj2.Empty()) //true
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type MyQueue struct {
	data  []int
	l     int
	first int
	last  int
}

func Constructor() MyQueue {
	return MyQueue{l: 0, data: []int{}, first: 0, last: 0}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
	this.l += 1
	this.last += 1
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	ans := this.data[this.first]
	this.data[this.first] = 0
	this.first += 1
	this.l -= 1
	return ans
}

func (this *MyQueue) Peek() int {
	if !this.Empty() {
		return this.data[this.first]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return this.l == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue2 struct {
	inStack, outStack []int
}

func Constructor2() MyQueue2 {
	return MyQueue2{}
}

func (q *MyQueue2) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue2) Pop() int {
	if len(q.outStack) == 0 {
		q.outStack = append(q.outStack, q.inStack...)
		q.inStack = []int{}
	}
	ans := q.outStack[0]
	q.outStack = q.outStack[1:]
	return ans
}

func (q *MyQueue2) Peek() int {
	if len(q.outStack) == 0 {
		q.outStack = append(q.outStack, q.inStack...)
		q.inStack = []int{}
	}
	return q.outStack[0]
}

func (q *MyQueue2) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
