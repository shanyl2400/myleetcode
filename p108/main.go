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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	node := &TreeNode{
		Val: nums[idx],
	}

	node.Left = sortedArrayToBST(nums[:len(nums)/2])
	node.Right = sortedArrayToBST(nums[len(nums)/2+1:])

	return node
}

func main() {
	tree := sortedArrayToBST([]int{
		-10, -3, 0, 5, 9,
	})
	fmt.Println(tree)
	fmt.Println(tree.Left)
	fmt.Println(tree.Left.Left)
	fmt.Println(tree.Right)
	fmt.Println(tree.Right.Left)
}
