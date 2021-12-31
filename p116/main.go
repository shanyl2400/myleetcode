package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type HeightNode struct {
	Height int
	Node   *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*HeightNode, 1)
	queue[0] = &HeightNode{
		Height: 1,
		Node:   root,
	}
	prev := &HeightNode{
		Height: 0,
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if prev.Height == node.Height {
			prev.Node.Next = node.Node
		}
		prev = node

		if node.Node.Left != nil {
			queue = append(queue, &HeightNode{
				Height: node.Height + 1,
				Node:   node.Node.Left,
			})
		}
		if node.Node.Right != nil {
			queue = append(queue, &HeightNode{
				Height: node.Height + 1,
				Node:   node.Node.Right,
			})
		}
	}

	return root
}

func tree0() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	return node1
}

func main() {
	l := connect(tree0())

	fmt.Println(l)
}
