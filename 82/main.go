package main

import "fmt"

func main() {
	node := ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							Next: &ListNode{
								Next: &ListNode{},
								Val: 5,
							},
							Val: 4,
						},
						Val: 4,
					},
					Val: 3,
				},
				Val: 3,
			}, Val: 2}, Val: 1}
	fmt.Println(deleteDuplicates(&node))
	// ch := make(chan int)
	// ch <- 2
	// fmt.Println(<-ch)
}

type ListNode struct {
	Val int
	Next *ListNode
}



func deleteDuplicates(head *ListNode) *ListNode {
	d := &ListNode{0, head}
	cur := d

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return d.Next
}
