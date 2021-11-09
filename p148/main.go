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

func swap(x, y, prevX, prevY *ListNode, begin **ListNode) {
	// fmt.Println("swap:", x, y, x.Next, y.Next)
	if x == y {
		return
	}
	if x == *begin {
		*begin = y
	} else if y == *begin {
		*begin = x
	}
	if prevX != nil {
		prevX.Next = y
	}
	if prevY != nil {
		prevY.Next = x
	}
	tmp := x.Next
	x.Next = y.Next
	y.Next = tmp
	// tmp := x.Val
	// x.Val = y.Val
	// y.Val = tmp
	// fmt.Println("aft swap:", y, x, x.Next, y.Next)

}

func quickSortOnce(begin *ListNode, end *ListNode, cnt int, name string) *ListNode {
	if begin == nil || end == nil || begin == end {
		return nil
	}

	var prevFast *ListNode
	var prevSlow *ListNode
	alreadyPrevFast := false
	alreadyPrevSlow := false
	fast := begin
	slow := begin
	leftCount := 1
	for fast != end {
		if fast.Val < end.Val {
			beforeBegin := begin
			fmt.Println("before slow:", slow, "fast:", fast, "prevSlow:", prevSlow, "prevFast:", prevFast)
			swap(slow, fast, prevSlow, prevFast, &begin)
			tmp := fast
			fast = slow
			slow = tmp
			fmt.Println("aft slow:", slow, "fast:", fast, "prevSlow:", prevSlow, "prevFast:", prevFast)
			if beforeBegin != begin {
				if prevFast == beforeBegin && !alreadyPrevFast {
					prevFast = begin
				}
				if prevSlow == beforeBegin && !alreadyPrevSlow {
					prevSlow = begin
				}
			}

			fmt.Println("switch >>>", fast, slow, fast.Next, slow.Next, prevFast, prevSlow)

			slow = slow.Next
			if prevSlow == nil {
				prevSlow = begin
			} else {
				alreadyPrevSlow = true
				prevSlow = prevSlow.Next
			}
			fmt.Println("next>>>>>>", fast, slow, fast.Next, slow.Next, prevFast, prevSlow)
			leftCount++
		}
		fast = fast.Next
		if prevFast == nil {
			prevFast = begin
		} else {
			alreadyPrevFast = true
			prevFast = prevFast.Next
		}
		fmt.Println("2222next>>>>>>", fast, slow, fast.Next, slow.Next, prevFast, prevSlow)
	}
	swap(slow, end, prevSlow, prevFast, &begin)
	if slow != begin {
		quickSortOnce(begin, prevSlow, leftCount, "left")
	}
	if slow != end {
		quickSortOnce(slow.Next, end, cnt-leftCount-1, "right")
	}
	return begin
}

func sortList(head *ListNode) *ListNode {
	pt := head
	cnt := 0
	for pt.Next != nil {
		pt = pt.Next
		cnt++
	}
	ret := quickSortOnce(head, pt, cnt, "mid")
	return ret
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

func main() {
	root := link1()
	ret := sortList(root)
	printLink(ret, nil)
	// printLink(root, nil)
	// node4 := root.Next
	// node1 := node4.Next
	// swap(node4, node1, root, node4, &root)
	// printLink(root, nil)
}
