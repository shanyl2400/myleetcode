package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := root.Val
	doMaxPathSum(root, &max)
	return max
}
func doMaxPathSum(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	leftMax := doMaxPathSum(root.Left, max)
	rightMax := doMaxPathSum(root.Right, max)
	ret := root.Val
	// fmt.Println("leftMax", leftMax, "rightMax", rightMax, "val", ret)
	if leftMax > rightMax {
		if leftMax > 0 {
			ret = ret + leftMax
		}
		if *max < ret+rightMax {
			*max = ret + rightMax
		}
	} else {
		if rightMax > 0 {
			ret = ret + rightMax
		}
		if *max < ret+leftMax {
			*max = ret + leftMax
		}
	}

	if *max < ret {
		*max = ret
	}
	return ret
}
func tree1() *TreeNode {
	node1 := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node1.Left = node2
	node1.Right = node3
	return node1
}

func tree2() *TreeNode {
	node_10 := &TreeNode{
		Val: -10,
	}
	node9 := &TreeNode{
		Val: 9,
	}
	node20 := &TreeNode{
		Val: 20,
	}
	node15 := &TreeNode{
		Val: 15,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node_10.Left = node9
	node_10.Right = node20
	node20.Left = node15
	node20.Right = node7
	return node_10
}

func tree3() *TreeNode {
	return &TreeNode{
		Val: 0,
	}
}

func tree4() *TreeNode {
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node8 := &TreeNode{Val: 8}
	node11 := &TreeNode{Val: 11}
	node13 := &TreeNode{Val: 13}
	node_4 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1}

	node5.Left = node4
	node5.Right = node8
	node4.Left = node11
	node11.Left = node7
	node11.Right = node2
	node8.Left = node13
	node8.Right = node_4
	node_4.Right = node1
	return node5
}
func tree5() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	node1.Right = node2
	node2.Right = node3
	node3.Right = node4
	node4.Right = node5
	return node1
}

func tree6() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 0}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 2}
	node6 := &TreeNode{Val: 0}
	node7 := &TreeNode{Val: -1}
	node8 := &TreeNode{Val: 0}
	node9 := &TreeNode{Val: 1}
	node10 := &TreeNode{Val: -1}
	node11 := &TreeNode{Val: 0}
	node12 := &TreeNode{Val: -1}
	node13 := &TreeNode{Val: 0}
	node14 := &TreeNode{Val: 1}
	node15 := &TreeNode{Val: 0}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node4.Left = node8
	node4.Right = node9
	node5.Left = node10
	node5.Right = node11
	node6.Left = node12
	node6.Right = node13
	node7.Left = node14
	node7.Right = node15
	return node1
}

func main() {
	fmt.Println(maxPathSum(tree6()))
}
