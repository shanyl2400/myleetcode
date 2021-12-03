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

func calcPathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == targetSum {
		ret++
	}
	ret += calcPathSum(root.Left, targetSum-root.Val)
	ret += calcPathSum(root.Right, targetSum-root.Val)
	return ret
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ret := calcPathSum(root, targetSum)
	ret += pathSum(root.Left, targetSum)
	ret += pathSum(root.Right, targetSum)
	return ret
}

func tree0() *TreeNode {
	root := &TreeNode{Val: 10}
	l := &TreeNode{Val: 5}
	r := &TreeNode{Val: -3}
	ll := &TreeNode{Val: 3}
	lr := &TreeNode{Val: 2}
	lll := &TreeNode{Val: 3}
	llr := &TreeNode{Val: -2}
	lrr := &TreeNode{Val: 1}
	rr := &TreeNode{Val: 11}

	root.Left = l
	root.Right = r
	l.Left = ll
	l.Right = lr
	ll.Left = lll
	ll.Right = llr
	lr.Right = lrr
	r.Right = rr
	return root
}

func tree1() *TreeNode {
	root := &TreeNode{Val: 5}
	l := &TreeNode{Val: 4}
	r := &TreeNode{Val: 8}
	ll := &TreeNode{Val: 11}
	lll := &TreeNode{Val: 7}
	llr := &TreeNode{Val: 2}
	rr := &TreeNode{Val: 4}
	rl := &TreeNode{Val: 13}
	rrl := &TreeNode{Val: 5}
	rrr := &TreeNode{Val: 1}

	root.Left = l
	root.Right = r
	l.Left = ll
	ll.Left = lll
	ll.Right = llr
	r.Right = rr
	r.Left = rl
	rr.Left = rrl
	rr.Right = rrr
	return root
}

func tree2() *TreeNode {
	root := &TreeNode{Val: 1}
	l := &TreeNode{Val: 2}
	ll := &TreeNode{Val: 3}
	lll := &TreeNode{Val: 4}
	llll := &TreeNode{Val: 5}
	root.Left = l
	l.Left = ll
	ll.Left = lll
	lll.Left = llll
	return root
}
func main() {
	// root := tree1()
	// target := 22
	root := tree2()
	target := 3
	fmt.Println(pathSum(root, target))
}
