package main

import "fmt"

func solveSudoku(board [][]byte) {
	nextStep := nextDot(board, -1)
	doSolveSudoku(board, nextStep)
}
func doSolveSudoku(board [][]byte, step int) bool {
	//outbound return true
	if step >= 81 {
		return true
	}
	i := step / 9
	j := step % 9

	//find candidates
	candidates := findCandidates(board, i, j)
	if len(candidates) < 1 {
		//no candidates
		//向上回溯
		return false
	}
	for c := range candidates {
		board[i][j] = candidates[c]
		nextStep := nextDot(board, step)
		ret := doSolveSudoku(board, nextStep)
		if ret {
			return true
		}
	}
	board[i][j] = '.'
	return false
}
func nextDot(board [][]byte, step int) int {
	step = step + 1
	if step >= 81 {
		return step
	}
	for board[step/9][step%9] != '.' {
		step = step + 1
		if step >= 81 {
			return step
		}
	}
	return step
}

func findCandidates(board [][]byte, i, j int) []byte {
	row := (i / 3) * 3
	col := (j / 3) * 3
	candidateMap := make([]bool, 10)
	for x := 0; x < 9; x++ {
		if board[i][x] != '.' {
			candidateMap[board[i][x]-48] = true
		}
	}
	for x := 0; x < 9; x++ {
		if board[x][j] != '.' {
			candidateMap[board[x][j]-48] = true
		}
	}
	for y := row; y < row+3; y++ {
		for x := col; x < col+3; x++ {
			if board[y][x] != '.' {
				candidateMap[board[y][x]-48] = true
			}
		}
	}
	candidates := make([]byte, 0, 9)
	for i := range candidateMap[1:] {
		if !candidateMap[i+1] {
			candidates = append(candidates, byte(i+1+48))
		}
	}
	return candidates
}

func printBoard(board [][]byte) {
	for i := range board {
		for j := range board[0] {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}
	fmt.Println("---------")
}
func main() {
	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}

	solveSudoku(board)
	for i := range board {
		for j := range board[0] {
			fmt.Print(string(board[i][j]))
		}
		fmt.Println()
	}
}
