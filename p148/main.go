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

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	//计算node数量
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	dummyHead := &ListNode{Next: head}
	for subLen := 1; subLen < cnt; subLen <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for j := 1; j < subLen && cur.Next != nil; j++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for j := 1; j < subLen && cur != nil && cur.Next != nil; j++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			fmt.Println(head1, head2)

			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

func merge(f, s *ListNode) *ListNode {
	dummyHead := &ListNode{}
	fs := f
	ss := s

	cur := dummyHead
	for fs != nil && ss != nil {
		if fs.Val > ss.Val {
			cur.Next = ss
			ss = ss.Next
		} else {
			cur.Next = fs
			fs = fs.Next
		}
		cur = cur.Next
	}

	//f未结束
	if fs != nil {
		cur.Next = fs
	}
	//s未结束
	if ss != nil {
		cur.Next = ss
	}
	// cur.Next = nil
	return dummyHead.Next
}

func link1() *ListNode {
	node4 := &ListNode{
		Val: 4,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node1 := &ListNode{
		Val: 1,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4.Next = node2
	node2.Next = node1
	node1.Next = node3
	return node4
}

func link2() *ListNode {
	node_1 := &ListNode{
		Val: -1,
	}
	node5 := &ListNode{
		Val: 5,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node0 := &ListNode{
		Val: 0,
	}
	node_1.Next = node5
	node5.Next = node3
	node3.Next = node4
	node4.Next = node0
	return node_1
}

func link3() *ListNode {
	node2 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node1 := &ListNode{
		Val: 1,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node2.Next = node4
	node4.Next = node1
	node1.Next = node3
	return node2
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
	// root := arrayToLink([]int{4, 1, 3, 2, 8, 5, 10, 6})
	root := arrayToLink([]int{-1, 5, 3, 4, 0})
	printLink(root, nil)
	head := sortList(root)

	printLink(head, nil)
}
