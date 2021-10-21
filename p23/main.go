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

func mergeKLists(lists []*ListNode) *ListNode {
	var cur, head *ListNode

	for !isEnd(lists) {
		min := 100000
		index := 0
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				index = i
			}
		}
		if head == nil {
			head = &ListNode{
				Val: lists[index].Val,
			}
			cur = head
		} else {
			cur.Next = &ListNode{
				Val: lists[index].Val,
			}
			cur = cur.Next
		}

		fmt.Println(index)
		lists[index] = lists[index].Next
	}
	return head
}

func isEnd(lists []*ListNode) bool {
	for i := range lists {
		if lists[i] != nil {
			return false
		}
	}
	return true
}

func main() {
	// m1 := &ListNode{
	// 	Val: 5,
	// }
	// m2 := &ListNode{
	// 	Val:  4,
	// 	Next: m1,
	// }
	// m3 := &ListNode{
	// 	Val:  1,
	// 	Next: m2,
	// }

	// n1 := &ListNode{
	// 	Val: 4,
	// }
	// n2 := &ListNode{
	// 	Val:  3,
	// 	Next: n1,
	// }
	// n3 := &ListNode{
	// 	Val:  1,
	// 	Next: n2,
	// }

	// p1 := &ListNode{
	// 	Val: 6,
	// }
	// p2 := &ListNode{
	// 	Val:  2,
	// 	Next: p1,
	// }

	out := mergeKLists([]*ListNode{})

	for out != nil {
		fmt.Print(out.Val, "->")
		out = out.Next
	}
}
