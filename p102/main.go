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
type Item struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []Item{
		{
			node:  root,
			level: 0,
		},
	}
	ret := make([][]int, 0)

	for len(queue) > 0 {
		//dequeue
		item := queue[0]
		queue = queue[1:]

		if len(ret) <= item.level {
			ret = append(ret, []int{
				item.node.Val,
			})
		} else {
			ret[item.level] = append(ret[item.level], item.node.Val)
		}
		if item.node.Left != nil {
			queue = append(queue, Item{
				node:  item.node.Left,
				level: item.level + 1,
			})
		}
		if item.node.Right != nil {
			queue = append(queue, Item{
				node:  item.node.Right,
				level: item.level + 1,
			})
		}

	}

	return ret
}

func main() {
	node3 := &TreeNode{
		Val: 3,
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

	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7

	fmt.Println(levelOrder(node3))
}
