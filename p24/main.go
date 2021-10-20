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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var first *ListNode
	second := head
	third := head.Next

	for third != nil {
		if first != nil {
			first.Next = third
		} else {
			head = third
		}
		second.Next = third.Next
		third.Next = second

		first = third.Next
		second = first.Next
		if second == nil {
			break
		}
		third = second.Next
	}
	return head
}

func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	// l4 := &ListNode{
	// 	Val:  4,
	// 	Next: l5,
	// }

	// l3 := &ListNode{
	// 	Val:  3,
	// 	Next: l4,
	// }

	// l2 := &ListNode{
	// 	Val:  2,
	// 	Next: l3,
	// }

	// l1 := &ListNode{
	// 	Val:  1,
	// 	Next: l2,
	// }
	h := swapPairs(l5)
	for h != nil {
		fmt.Print(h.Val, "->")
		h = h.Next
	}
	fmt.Println()
}
