package main

import "fmt"

type treeNode struct {
	children [26]*treeNode
	word     string
}

var seen map[string]bool
var dirs = []struct{ row, col int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (t *treeNode) insert(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &treeNode{}
		}
		node = node.children[ch]
	}
	node.word = word
}

func dfs(board [][]byte, t *treeNode, row, col int) {
	ch := board[row][col]
	node := t.children[ch-'a']
	if node == nil {
		return
	}
	if node.word != "" {
		seen[node.word] = true
	}
	board[row][col] = '#'
	for i := range dirs {
		newRow, newCol := row+dirs[i].row, col+dirs[i].col
		if newRow >= 0 && newRow < len(board) && newCol >= 0 &&
			newCol < len(board[0]) &&
			board[newRow][newCol] != '#' {
			dfs(board, node, newRow, newCol)
		}
	}
	board[row][col] = ch
}

func findWords(board [][]byte, words []string) []string {
	seen = make(map[string]bool)
	t := new(treeNode)
	for i := range words {
		t.insert(words[i])
	}
	for row := range board {
		for col := range board[row] {
			dfs(board, t, row, col)
		}
	}
	ans := make([]string, 0)
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oathfiie", "pea", "eat", "rain"}

	// board := [][]byte{
	// 	{'a', 'b'},
	// 	{'c', 'd'},
	// }
	// words := []string{"abcd"}
	fmt.Println(findWords(board, words))
}
