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

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first := head
	second := first.Next
	flag := false
	for second != nil {
		if first == second {
			flag = true
			break
		}
		first = first.Next
		second = second.Next
		if second == nil {
			return nil
		}
		second = second.Next
	}
	if flag {
		//record count
		first = first.Next
		cnt := 1
		for first != second {
			first = first.Next
			cnt++
		}
		//ring count is cnt
		fmt.Println(cnt)
		first = head
		second = head
		for i := 0; i < cnt; i++ {
			second = second.Next
		}
		for first != second {
			first = first.Next
			second = second.Next
		}
		return first
	}
	return nil
}
