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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rt := preorder[0]
	idx := 0
	for i := range inorder {
		if rt == inorder[i] {
			idx = i
			break
		}
	}
	node := &TreeNode{
		Val: rt,
	}

	leftInorder := inorder[:idx]
	rightInorder := inorder[idx+1:]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[1+len(leftInorder):]
	// fmt.Println("inorder:", leftInorder, rightInorder)
	// fmt.Println("preorder:", leftPreorder, rightPreorder)
	if len(leftInorder) > 0 {
		node.Left = buildTree(leftPreorder, leftInorder)
	}
	if len(rightInorder) > 0 {
		node.Right = buildTree(rightPreorder, rightInorder)
	}

	return node
}

func main() {
	preorder := []int{
		3, 9, 20, 15, 7,
	}
	inorder := []int{
		9, 3, 15, 20, 7,
	}
	root := buildTree(preorder, inorder)

	fmt.Println(root)
}
