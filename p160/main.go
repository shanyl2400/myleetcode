package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	top1 := -1
	top2 := -1

	ha, hb := headA, headB
	for ha != nil || hb != nil {
		if ha != nil {
			top1++
			stack1 = append(stack1, ha)
			ha = ha.Next
		}
		if hb != nil {
			top2++
			stack2 = append(stack2, hb)
			hb = hb.Next
		}
	}
	var out *ListNode
	for top1 >= 0 && top2 >= 0 {
		if stack1[top1] != stack2[top2] {
			break
		}
		out = stack1[top1]
		top1--
		top2--
	}
	return out
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

func findTail(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	for node.Next != nil {
		node = node.Next
	}
	return node
}
func main() {
	head1 := arrayToLink([]int{4, 1})
	head2 := arrayToLink([]int{5, 0, 1})
	tail := arrayToLink([]int{8, 4, 5})

	tail1 := findTail(head1)
	tail2 := findTail(head2)
	tail1.Next = tail
	tail2.Next = tail

	printLink(head1, nil)
	printLink(head2, nil)

	common := getIntersectionNode(head1, head2)
	printLink(common, nil)
}
