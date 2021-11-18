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

func isPalindrome(head *ListNode) bool {
	front := head

	var recursiveCheck func(cur *ListNode) bool
	recursiveCheck = func(cur *ListNode) bool {
		if cur != nil {
			if !recursiveCheck(cur.Next) {
				return false
			}
			if cur.Val != front.Val {
				return false
			}
			front = front.Next
		}
		return true
	}

	return recursiveCheck(head)
}
func isPalindrome2(head *ListNode) bool {
	stack := make([]int, 0)
	top := -1

	cur := head
	for cur != nil {
		top++
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		val := stack[top]
		if val != cur.Val {
			return false
		}
		top--

		cur = cur.Next
	}

	return true
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

func main() {
	ret := isPalindrome(arrayToLink([]int{
		// 1, 2, 2, 1,
		1, 2,
	}))
	fmt.Println(ret)
}
