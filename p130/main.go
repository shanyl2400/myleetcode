package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	size := len(board) * len(board[0])
	groups := make([]int, size)
	maxGrp := 1
	for i := range board {
		for j := range board[i] {
			idx := i*len(board[i]) + j
			if board[i][j] == 'O' && groups[idx] == 0 {
				//dfs
				dfs(board, idx, groups, maxGrp, make([]bool, size))
				maxGrp++
			}
		}
	}
	retains := make(map[int]bool)
	for i := range groups {
		d1 := i / len(board[0])
		d2 := i % len(board[0])
		if (d1 == 0 || d1 == len(board)-1 || d2 == 0 || d2 == len(board[0])-1) && groups[i] != 0 {
			//board
			retains[groups[i]] = true
		}
	}

	for i := range groups {
		if groups[i] > 0 && !retains[groups[i]] {
			//清零
			board[i/len(board[0])][i%len(board[0])] = 'X'
		}
	}
}

func dfs(board [][]byte, idx int, groups []int, grp int, visited []bool) {
	if idx < 0 || idx >= len(groups) || board[idx/len(board[0])][idx%len(board[0])] != 'O' {
		return
	}

	if !visited[idx] {
		groups[idx] = grp
		visited[idx] = true

		//around
		if idx-1 >= 0 && idx-1 < len(groups) {
			dfs(board, idx-1, groups, grp, visited)
		}
		if idx+1 >= 0 && idx+1 < len(groups) {
			dfs(board, idx+1, groups, grp, visited)
		}

		if idx-len(board[0]) >= 0 && idx-len(board[0]) < len(groups) {
			dfs(board, idx-len(board[0]), groups, grp, visited)
		}

		if idx+len(board[0]) >= 0 && idx+len(board[0]) < len(groups) {
			dfs(board, idx+len(board[0]), groups, grp, visited)
		}
	}
}

func main() {
	board := [][]byte{
		// {'X', 'X', 'X', 'X'},
		// {'X', 'O', 'O', 'X'},
		// {'X', 'X', 'O', 'X'},
		// {'X', 'O', 'X', 'X'},
		{'X'},
	}
	solve(board)
	for i := range board {
		for j := range board {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}
