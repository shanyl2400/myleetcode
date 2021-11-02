package main

import "fmt"

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := range board {
		for j := range board[i] {
			flag := doCheck(board, visited, i, j, 0, word)
			if flag {
				return true
			}
		}
	}

	return false
}
func doCheck(board [][]byte, visited [][]bool, i, j, idx int, word string) bool {
	//非要求字母，直接抛弃
	if board[i][j] != word[idx] {
		return false
	}
	//处理完成
	if idx >= len(word)-1 {
		return true
	}

	//标记visited
	visited[i][j] = true
	result := false
	for x := range dirs {
		ni := dirs[x][0] + i
		nj := dirs[x][1] + j
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[i]) && !visited[ni][nj] {
			//在区间内
			result = doCheck(board, visited, ni, nj, idx+1, word)
			if result {
				break
			}
		}
	}
	//回溯
	visited[i][j] = false
	return result
}

func main() {
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCCEA"
	// word := "ABCCED"
	// word := "ABCB"
	// word := "ABC"

	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'E', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	// word := "ABCESEEEFS"
	board := [][]byte{
		{'a'},
	}
	word := "a"
	fmt.Println(exist(board, word))
}
