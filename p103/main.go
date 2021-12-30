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
type TreeHeightNode struct {
	Node   *TreeNode
	Height int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeHeightNode, 1)
	queue[0] = &TreeHeightNode{
		Node:   root,
		Height: 1,
	}
	ret := make([][]int, 0)
	curHeight := 1
	curArray := make([]int, 0)
	for len(queue) > 0 {
		data := queue[0]
		queue = queue[1:]

		if data.Height > curHeight {
			ret = append(ret, curArray)
			curArray = make([]int, 0)
			curHeight = data.Height
		}
		curArray = append(curArray, data.Node.Val)

		if data.Node.Left != nil {
			queue = append(queue, &TreeHeightNode{
				Node:   data.Node.Left,
				Height: data.Height + 1,
			})
		}
		if data.Node.Right != nil {
			queue = append(queue, &TreeHeightNode{
				Node:   data.Node.Right,
				Height: data.Height + 1,
			})
		}
	}
	ret = append(ret, curArray)

	//odd array
	for i := 1; i < len(ret); i += 2 {
		//resort
		for j := 0; j < len(ret[i])/2; j++ {
			ret[i][j], ret[i][len(ret[i])-1-j] = ret[i][len(ret[i])-1-j], ret[i][j]
		}
	}

	return ret
}

func tree0() *TreeNode {
	node3 := &TreeNode{Val: 3}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node15 := &TreeNode{Val: 15}
	node7 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node10 := &TreeNode{Val: 10}

	node3.Left = node9
	node3.Right = node20
	node9.Left = node6
	node9.Right = node10
	node20.Left = node15
	node20.Right = node7
	return node3
}
func main() {
	fmt.Println(zigzagLevelOrder(tree0()))
}
