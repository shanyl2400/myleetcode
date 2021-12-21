package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}

func tree1() *TreeNode {
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 8}
	node0 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node9 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}

	node6.Left = node2
	node6.Right = node8
	node2.Left = node0
	node2.Right = node4
	node4.Left = node3
	node4.Right = node5

	node8.Left = node7
	node8.Right = node9

	return node6
}

func main() {
	root := tree1()
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right).Val)
}
