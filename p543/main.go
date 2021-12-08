package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func doDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := doDiameter(root.Left)
	r := doDiameter(root.Right)

	candidate := l + r + 1
	if max < candidate {
		max = candidate
	}
	if l < r {
		l = r
	}
	return l + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	doDiameter(root)

	return max - 1
}

func tree0() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	return node1
}
func tree1() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}

	node1.Right = node2
	return node1
}

func main() {
	fmt.Println(diameterOfBinaryTree(tree1()))
}
