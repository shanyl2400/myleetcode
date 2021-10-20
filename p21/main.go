package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	t1 := l1
	t2 := l2
	var head *ListNode
	var pointer *ListNode
	for t1 != nil && t2 != nil {
		node := &ListNode{}
		if t1.Val < t2.Val {
			node.Val = t1.Val
			t1 = t1.Next
		} else {
			node.Val = t2.Val
			t2 = t2.Next
		}
		if head == nil {
			head = node
			pointer = head
		} else {
			pointer.Next = node
			pointer = pointer.Next
		}
	}
	for t1 != nil {
		node := &ListNode{
			Val: t1.Val,
		}
		if head == nil {
			head = node
			pointer = head
		} else {
			pointer.Next = node
			pointer = pointer.Next
		}
		t1 = t1.Next
	}
	for t2 != nil {
		node := &ListNode{
			Val: t2.Val,
		}
		if head == nil {
			head = node
			pointer = head
		} else {
			pointer.Next = node
			pointer = pointer.Next
		}
		t2 = t2.Next
	}
	return head
}

func main() {
	l13 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l12 := &ListNode{
		Val:  3,
		Next: l13,
	}

	l11 := &ListNode{
		Val:  1,
		Next: l12,
	}

	l23 := &ListNode{
		Val:  4,
		Next: nil,
	}

	l22 := &ListNode{
		Val:  2,
		Next: l23,
	}
	l21 := &ListNode{
		Val:  1,
		Next: l22,
	}
	// for l11 != nil {
	// 	fmt.Print(l11.Val, "->")
	// 	l11 = l11.Next
	// }
	// fmt.Println()
	// for l21 != nil {
	// 	fmt.Print(l21.Val, "->")
	// 	l21 = l21.Next
	// }
	// fmt.Println()

	h := mergeTwoLists(l11, l21)
	for h != nil {
		fmt.Print(h.Val, "->")
		h = h.Next
	}
	fmt.Println()
}
