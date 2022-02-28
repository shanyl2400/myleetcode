package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	evenHead := head.Next
	odd := head
	even := odd.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func newLink(size int) *ListNode {
	var head *ListNode
	var cur *ListNode
	for i := 0; i < size; i++ {
		if i == 0 {
			head = &ListNode{Val: i + 1}
			cur = head
		} else {
			cur.Next = &ListNode{Val: i + 1}
			cur = cur.Next
		}
	}
	return head
}

func main() {
	head := newLink(5)
	head = oddEvenList(head)

	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()
}
