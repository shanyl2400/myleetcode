package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightest(root *TreeNode) *TreeNode {
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//调整左边
	right := root.Right
	if root.Left != nil {
		flatten(root.Left)
		rightestLeft := rightest(root.Left)
		oldRight := root.Right
		root.Right = root.Left
		root.Left = nil
		rightestLeft.Right = oldRight
	}
	flatten(right)

}

func tree1() *TreeNode {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node5 := &TreeNode{
		Val: 5,
	}
	node6 := &TreeNode{
		Val: 6,
	}
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Right = node6

	return node1
}
func tree2() *TreeNode {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node1.Right = node2
	node2.Left = node3

	return node1
}

func main() {
	root := tree2()
	flatten(root)
	cur := root
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Right
	}
	fmt.Println()
}
