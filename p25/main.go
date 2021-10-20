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

func reverseKGroup(head *ListNode, k int) *ListNode {
	first := head
	last := head
	nextFirst := head

	lastStack := make([]*ListNode, 0)
	for first != nil {
		lastStack = append(lastStack, first)
		first = first.Next
	}
	base := 0
	first = head

	for first != nil {
		//移动last
		for i := 0; i < k-1; i++ {
			last = last.Next
			if last == nil {
				return head
			}
		}

		nextFirst = last.Next
		//交换last和first
		lastIndex := base + k - 1
		last = lastStack[lastIndex]
		for i := 0; i < k/2; i++ {
			temp := last.Val
			last.Val = first.Val
			first.Val = temp

			// last --
			lastIndex--
			last = lastStack[lastIndex]

			// first ++
			first = first.Next
		}

		first = nextFirst
		last = nextFirst
		base = base + k
	}
	return head
}

func print(n *ListNode) {
	for n != nil {
		fmt.Print(n.Val, "->")
		n = n.Next
	}
}

func main() {
	p1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	p2 := &ListNode{
		Val:  4,
		Next: p1,
	}
	p3 := &ListNode{
		Val:  3,
		Next: p2,
	}
	p4 := &ListNode{
		Val:  2,
		Next: p3,
	}
	p5 := &ListNode{
		Val:  1,
		Next: p4,
	}

	head := reverseKGroup(p5, 2)
	print(head)
}
