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

func inOrder(root *TreeNode, f func(node *TreeNode)) {
	if root == nil {
		return
	}
	inOrder(root.Left, f)
	f(root)
	inOrder(root.Right, f)
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	idx := 1
	ret := -1
	inOrder(root, func(node *TreeNode) {
		if idx == k {
			ret = node.Val
		}
		idx++
	})
	return ret
}

func tree1() *TreeNode {
	node3 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}

	node3.Left = node1
	node1.Right = node2
	node3.Right = node4

	return node3
}
func tree2() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	node5.Left = node3
	node5.Right = node6
	node3.Left = node2
	node3.Right = node4
	node2.Left = node1
	return node5
}

func main() {
	fmt.Println(kthSmallest(tree2(), 3))
}
