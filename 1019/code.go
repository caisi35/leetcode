package main

type ListNode struct {
	Val int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	var ans []int
	var stack [][]int
	cur := head
	idx := -1
	for cur != nil {
		idx++
		ans = append(ans, 0)
		for len(stack) > 0 && stack[len(stack)-1][0] < cur.Val {
			top := stack[len(stack)-1]
			stack =stack[:len(stack)-1]
			ans[top[1]] = cur.Val
		}
		stack = append(stack, []int{cur.Val, idx})
	}
	return ans
}

func main() {
	list_node := ListNode{Val: 2}
	nextLargerNodes(&list_node)
}