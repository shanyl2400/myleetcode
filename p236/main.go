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

var trace1 []*TreeNode
var trace2 []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	trace1 = nil
	trace2 = nil
	if root == nil || p == nil || q == nil {
		return nil
	}
	res := findChild(root, p, q, []*TreeNode{})
	if !res {
		return nil
	}
	ret := root
	for i := 0; i < len(trace1) && i < len(trace2) && trace1[i] == trace2[i]; i++ {
		if trace1[i] == trace2[i] {
			ret = trace1[i]
		}
	}

	return ret
}

func findChild(root, p, q *TreeNode, trace []*TreeNode) bool {
	trace = append(trace, root)
	if root == p {
		trace1 = trace
		if trace2 != nil {
			return true
		}
	}
	if root == q {
		trace2 = trace
		if trace1 != nil {
			return true
		}
	}
	if root.Left != nil {
		ret := findChild(root.Left, p, q, trace)
		if ret {
			return true
		}
	}
	if root.Right != nil {
		ret := findChild(root.Right, p, q, trace)
		if ret {
			return true
		}
	}
	return false
}
func tree1() *TreeNode {
	node3 := &TreeNode{
		Val: 3,
	}
	node5 := &TreeNode{
		Val: 5,
	}
	node1 := &TreeNode{
		Val: 1,
	}
	node6 := &TreeNode{
		Val: 6,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node0 := &TreeNode{
		Val: 0,
	}
	node8 := &TreeNode{
		Val: 8,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node3.Left = node5
	node3.Right = node1
	node5.Left = node6
	node5.Right = node2
	node1.Left = node0
	node1.Right = node8
	node2.Left = node7
	node2.Right = node4

	return node3
}

func main() {
	root := tree1()
	p := root.Left
	// q := p.Right.Right
	q := root.Right
	ret := lowestCommonAncestor(root, p, q)
	fmt.Println(ret.Val)
}
