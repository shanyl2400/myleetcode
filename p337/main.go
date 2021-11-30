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

func max(x, y int) int {
	if x < y {
		x = y
	}
	return x
}
func childMax(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return rob(node.Left) + rob(node.Right)
}
func maxCMax(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, 0
	}
	lmax, lcmax, rmax, rcmax := 0, 0, 0, 0
	if node.Left != nil {
		lmax, lcmax = maxCMax(node.Left)
	}
	if node.Right != nil {
		rmax, rcmax = maxCMax(node.Right)
	}
	cmax := max(lmax, lcmax) + max(rcmax, rmax)
	max := lcmax + rcmax + node.Val
	return max, cmax
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// val1 := rob(root.Left) + rob(root.Right)
	// val2 := root.Val + childMax(root.Left) + childMax(root.Right)
	// return max(val1, val2)
	return max(maxCMax(root))
}

func tree1() *TreeNode {
	root := &TreeNode{Val: 3}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	lr := &TreeNode{Val: 3}
	rr := &TreeNode{Val: 1}

	root.Left = l
	root.Right = r
	l.Right = lr
	r.Right = rr
	return root
}

func tree2() *TreeNode {
	root := &TreeNode{Val: 3}
	l := &TreeNode{Val: 4}
	r := &TreeNode{Val: 5}
	ll := &TreeNode{Val: 1}
	lr := &TreeNode{Val: 3}
	rr := &TreeNode{Val: 1}

	root.Left = l
	root.Right = r
	l.Left = ll
	l.Right = lr
	r.Right = rr
	return root
}

func tree3() *TreeNode {
	root := &TreeNode{Val: 4}
	l := &TreeNode{Val: 1}
	ll := &TreeNode{Val: 2}
	lll := &TreeNode{Val: 3}

	root.Left = l
	l.Left = ll
	ll.Left = lll
	return root
}

func tree4() *TreeNode {
	root := &TreeNode{Val: 2}
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 3}
	lr := &TreeNode{Val: 4}

	root.Left = l
	root.Right = r
	l.Right = lr
	return root
}
func main() {
	fmt.Println(rob(tree4()))
}
