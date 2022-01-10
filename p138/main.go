package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = make(map[*Node]*Node)
	return deepCopy(head)
}
func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func randLink1() *Node {
	node7 := &Node{
		Val: 7,
	}
	node13 := &Node{
		Val: 13,
	}
	node11 := &Node{
		Val: 11,
	}
	node10 := &Node{
		Val: 10,
	}
	node1 := &Node{
		Val: 1,
	}

	node7.Next = node13
	node13.Next = node11
	node11.Next = node10
	node10.Next = node1

	node13.Random = node7
	node11.Random = node1
	node10.Random = node11
	node1.Random = node1

	return node7
}

func main() {
	link1 := randLink1()
	link2 := copyRandomList(link1)

	for link2 != nil {
		fmt.Print(link2.Val)
		if link2.Random != nil {
			fmt.Printf("(%v)", link2.Random.Val)
		}
		fmt.Print("->")
		link2 = link2.Next
	}
	fmt.Println()
}
