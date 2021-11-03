package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	//中序遍历
	return doIsSymmetric(root.Left, root.Right)
}

func doIsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left.Val == right.Val {
		flag := doIsSymmetric(left.Left, right.Right)
		if !flag {
			return false
		}
		return doIsSymmetric(left.Right, right.Left)
	}

	return false
}

func tree1() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	l := &TreeNode{
		Val: 2,
	}
	r := &TreeNode{
		Val: 2,
	}
	ll := &TreeNode{
		Val: 3,
	}

	lr := &TreeNode{
		Val: 4,
	}

	rl := &TreeNode{
		Val: 4,
	}
	rr := &TreeNode{
		Val: 3,
	}

	root.Left = l
	root.Right = r
	l.Left = ll
	l.Right = lr
	r.Left = rl
	r.Right = rr
	return root
}

func tree2() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	l := &TreeNode{
		Val: 2,
	}
	r := &TreeNode{
		Val: 2,
	}
	ll := &TreeNode{
		Val: 2,
	}

	rl := &TreeNode{
		Val: 2,
	}

	root.Left = l
	root.Right = r
	l.Left = ll
	r.Left = rl
	return root
}
func main() {

	fmt.Println(isSymmetric(tree1()))
}
