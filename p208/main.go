package main

import "fmt"

type Trie struct {
	n        byte
	h        int
	children []*Trie
}

func Constructor() Trie {
	return Trie{
		children: make([]*Trie, 0),
	}
}

func (this *Trie) Insert(word string) {
	word = word + "."
	idx := 0
	flag := false
	cur := this
	x := 0
	for i := idx; i < len(word); i++ {
		flag = false
		for x := range cur.children {
			if cur.children[x].n == word[i] {
				cur = cur.children[x]
				idx++
				flag = true
				break
			}
		}
		if !flag {
			break
		}
		x++
	}

	// cur = this
	for i := idx; i < len(word); i++ {
		node := &Trie{
			n: word[i],
			h: cur.h + 1,
		}
		cur.children = append(cur.children, node)
		cur = node
	}
}

func (this *Trie) Search(word string) bool {
	word = word
	cur := this.pos(word)
	if cur == nil {
		return false
	}
	for i := range cur.children {
		if cur.children[i].n == '.' {
			return true
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.pos(prefix)
	if cur == nil {
		return false
	}

	return true
}

func (this *Trie) pos(word string) *Trie {
	cur := this
	flag := false
	for i := 0; i < len(word); i++ {
		flag = false
		for x := range cur.children {
			if cur.children[x].n == word[i] {
				cur = cur.children[x]
				flag = true
				break
			}
		}
		if !flag {
			return nil
		}
	}
	return cur
}

func (this *Trie) print() {
	queue := make([]*Trie, 0)
	queue = append(queue, this.children...)
	h := 0
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		fmt.Print(string(q.n), " ")
		if q.h > h {
			fmt.Println()
			h++
		}
		queue = append(queue, q.children...)
	}
	fmt.Println()
}
func main() {
	word := "apple"
	prefix := "app"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
	param_4 := obj.Search(prefix)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)

	obj.Insert(prefix)
	// obj.print()
	param_5 := obj.Search(prefix)
	fmt.Println(param_5)
}
