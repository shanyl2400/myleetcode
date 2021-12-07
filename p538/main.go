package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cur int

func convertBST(root *TreeNode) *TreeNode {
	//反中序遍历
	if root == nil {
		return nil
	}
	convertBST(root.Right)
	cur += root.Val
	root.Val = cur
	fmt.Println(cur)
	convertBST(root.Left)
	return root
}

func print(root *TreeNode) {
	doPrint(root)
	fmt.Println()
}
func doPrint(root *TreeNode) {
	if root == nil {
		return
	}
	doPrint(root.Right)
	fmt.Print(root.Val, "->")
	doPrint(root.Left)
}

func tree0() *TreeNode {
	node0 := &TreeNode{Val: 0}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 8}
	node4.Left = node1
	node4.Right = node6
	node1.Left = node0
	node1.Right = node2
	node2.Right = node3
	node6.Left = node5
	node6.Right = node7
	node7.Right = node8
	return node4
}

func main() {
	root := tree0()
	root = convertBST(root)
	print(root)
}
