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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	//calculate k
	ptr := head
	size := 0
	for ptr != nil {
		ptr = ptr.Next
		size++
	}
	k = k % size
	if k == 0 {
		return head
	}

	//find head
	ptr = head
	ptr2 := head
	for i := 0; i < k; i++ {
		ptr2 = ptr2.Next
	}
	for ptr2.Next != nil {
		ptr = ptr.Next
		ptr2 = ptr2.Next
	}

	//new head
	newHead := ptr.Next

	ptr.Next = nil

	ptr = newHead
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = head

	return newHead
}

func list1() *ListNode {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	return node1
}

func list2() *ListNode {
	node1 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2}

	node1.Next = node2
	node2.Next = node3
	return node1
}

func main() {
	root := list2()
	root2 := rotateRight(root, 4)

	for root2 != nil {
		fmt.Print(root2.Val, "->")
		root2 = root2.Next
	}
	fmt.Println()
}
