package main

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

func helper(root1 *TreeNode, root2 *TreeNode) {
	if root1 == nil || root2 == nil {
		if root2 != nil {
			root1 = root2
		}
		return
	}
	root1.Val += root2.Val
	if root1.Left != nil && root2.Left != nil {
		helper(root1.Left, root2.Left)
	}
	if root1.Right != nil && root2.Right != nil {
		helper(root1.Right, root2.Right)
	}

	if root1.Left == nil {
		root1.Left = root2.Left
	}
	if root1.Right == nil {
		root1.Right = root2.Right
	}
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	helper(root1, root2)
	if root2 != nil && root1 == nil {
		return root2
	}
	return root1
}
