package main

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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	first := head
	second := head.Next
	for second != nil {
		if first == second {
			return true
		}
		first = first.Next
		second = second.Next
		if second == nil {
			return false
		}
		second = second.Next
	}
	return false
}
