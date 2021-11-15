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

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := make([]*ListNode, 0)
	top := -1

	dummyHead := &ListNode{Next: head}
	cur := head
	for cur != nil {
		top++
		stack = append(stack, cur)
		cur = cur.Next
	}

	front := head
	preFront := dummyHead
	for i := 0; i < len(stack)/2; i++ {
		back := stack[top]
		top--
		preBack := stack[top]

		preFront.Next = back
		preBack.Next = front

		tmp := front.Next
		front.Next = back.Next
		back.Next = tmp

		front = back.Next
		preFront = back
	}

	return dummyHead.Next
}

func arrayToLink(a []int) *ListNode {
	root := &ListNode{Val: a[0]}
	cur := root
	for i := range a[1:] {
		cur.Next = &ListNode{
			Val: a[i+1],
		}
		cur = cur.Next
	}
	return root
}
func printLink(root, end *ListNode) {
	cur := root
	for cur != end {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	if end != nil {
		fmt.Print(cur.Val)
	}
	fmt.Println()
}
func main() {
	link := []int{1, 2, 3, 4, 5}
	// link := []int{1, 2}
	// link := []int{1}
	root := arrayToLink(link)
	printLink(root, nil)
	head := reverseList(root)

	printLink(head, nil)
}
