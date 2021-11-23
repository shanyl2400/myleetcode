package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

type TreeNodeParentIdx struct {
	Node   *TreeNode
	Parent int
	LR     string
}

type CodecItem struct {
	Val    int
	Parent int
	LR     string
}

func (c *CodecItem) String() string {
	return fmt.Sprintf("%d,%d,%v", c.Val, c.Parent, c.LR)
}
func Parse(d string) *CodecItem {
	items := strings.Split(d, ",")
	if len(items) != 3 {
		return nil
	}
	val, _ := strconv.Atoi(items[0])
	parent, _ := strconv.Atoi(items[1])
	lr := items[2]
	return &CodecItem{
		Val:    val,
		Parent: parent,
		LR:     lr,
	}
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	items := make([]*CodecItem, 0)
	idx := -1
	this.iterator(root, func(node *TreeNodeParentIdx) int {
		items = append(items, &CodecItem{
			Val:    node.Node.Val,
			Parent: node.Parent,
			LR:     node.LR,
		})
		idx++
		return idx
	})
	ret := ""
	for i := range items {
		ret = ret + items[i].String() + ";"
	}
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, ";")
	nodes := make([]*TreeNode, 0)
	for i := range tokens {
		item := Parse(tokens[i])
		if item == nil {
			continue
		}
		node := &TreeNode{
			Val: item.Val,
		}
		nodes = append(nodes, node)
		if item.Parent != -1 {
			if item.LR == "l" {
				nodes[item.Parent].Left = node
			} else {
				nodes[item.Parent].Right = node
			}
		}

	}
	return nodes[0]
}

func (this *Codec) iterator(root *TreeNode, handle func(node *TreeNodeParentIdx) int) {
	queue := make([]*TreeNodeParentIdx, 0)
	queue = append(queue, &TreeNodeParentIdx{
		Node:   root,
		Parent: -1,
	})
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		idx := handle(q)
		if q.Node.Left != nil {
			queue = append(queue, &TreeNodeParentIdx{
				Node:   q.Node.Left,
				Parent: idx,
				LR:     "l",
			})
		}
		if q.Node.Right != nil {
			queue = append(queue, &TreeNodeParentIdx{
				Node:   q.Node.Right,
				Parent: idx,
				LR:     "r",
			})
		}
	}
}

func tree1() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	// node8 := &TreeNode{Val: 8}
	node9 := &TreeNode{Val: 9}
	node10 := &TreeNode{Val: 10}
	node11 := &TreeNode{Val: 11}
	node12 := &TreeNode{Val: 12}
	node13 := &TreeNode{Val: 13}
	node14 := &TreeNode{Val: 14}
	node15 := &TreeNode{Val: 15}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	// node4.Left = node8
	node4.Right = node9
	node5.Left = node10
	node5.Right = node11
	node6.Left = node12
	node6.Right = node13
	node7.Left = node14
	node7.Right = node15
	return node1
}

func tree2() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node1.Right = node2

	return node1
}
func tree3() *TreeNode {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	node4.Left = node6
	node4.Right = node7

	return node1
}

func tree4() *TreeNode {
	var root *TreeNode
	var cur *TreeNode
	for i := 1; i <= 1000; i++ {
		node := &TreeNode{Val: i}
		if i == 1 {
			root = node
			cur = node
		}
		cur.Right = node
		cur = node
	}
	return root
}
func main() {
	// root := tree4()
	var root *TreeNode
	codec := Constructor()
	data := codec.serialize(root)
	fmt.Println(data)
	root2 := codec.deserialize(data)
	codec.iterator(root2, func(node *TreeNodeParentIdx) int {
		if node.Node != nil {
			fmt.Print(node.Node.Val, "->")
		}
		return 0
	})
	fmt.Println()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
