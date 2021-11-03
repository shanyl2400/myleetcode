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

func findVal(root *TreeNode, left bool) int {
	cur := root
	ret := cur.Val
	for cur != nil {
		ret = cur.Val
		if left {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ret
}

func isValidBST(root *TreeNode) bool {
	//left
	if root.Left != nil {
		//find left max
		val := findVal(root.Left, false)
		if val >= root.Val {
			return false
		}
		flag := isValidBST(root.Left)
		if !flag {
			return false
		}
	}

	//right
	if root.Right != nil {
		//find right min
		val := findVal(root.Right, true)
		if val <= root.Val {
			return false
		}
		flag := isValidBST(root.Right)
		if !flag {
			return false
		}
	}
	return true
}

func main() {
	node5 := TreeNode{
		Val: 5,
	}
	node4 := TreeNode{
		Val: 4,
	}
	node6 := TreeNode{
		Val: 7,
	}
	node7 := TreeNode{
		Val: 8,
	}
	node3 := TreeNode{
		Val: 6,
	}
	node5.Left = &node4
	node5.Right = &node6
	node6.Left = &node3
	node6.Right = &node7
	fmt.Println(isValidBST(&node5))
}
