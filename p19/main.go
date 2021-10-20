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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	heap := make([]*ListNode, 0)
	p := head
	for p != nil {
		heap = append(heap, p)
		p = p.Next
	}
	index := len(heap) - n
	if index < 0 {
		return head
	}
	var nextPointer *ListNode
	if index+1 < len(heap) {
		nextPointer = heap[index+1]
	}
	if index-1 < 0 {
		head = nextPointer
	} else {
		heap[index-1].Next = nextPointer
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

	h := removeNthFromEnd(l5, 1)
	for h != nil {
		fmt.Print(h.Val, "->")
		h = h.Next
	}
	fmt.Println()
}
