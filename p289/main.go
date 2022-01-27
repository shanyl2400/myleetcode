package main

import "fmt"

var dirs = [][]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1},
}
var nextDirs = []int{
	0, -1, 1,
}

func calc(board [][]int, i, j int) int {
	sum := 0
	for _, dir := range dirs {
		ni := i + dir[0]
		nj := j + dir[1]
		if ni < 0 || ni >= len(board) || nj < 0 || nj >= len(board[ni]) {
			continue
		}
		sum += board[ni][nj] % 2
	}
	return sum
}

func calcNext(board [][]int, i, j int, prev int) int {
	sum := prev
	for _, dir := range nextDirs {
		ni := i + dir

		nj := j + 1
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[ni]) {
			sum += board[ni][nj] % 2
		}

		nj = j - 2
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[ni]) {
			sum -= board[ni][nj] % 2
		}
	}
	//add & remove it self
	sum -= board[i][j] % 2
	if j-1 >= 0 {
		sum += board[i][j-1] % 2
	}
	return sum
}

func calcRow(board [][]int, i, j int, prev int) int {
	sum := prev
	for _, dir := range nextDirs {
		nj := j + dir

		ni := i + 1
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[ni]) {
			sum += board[ni][nj] % 2
		}

		ni = i - 2
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[ni]) {
			sum -= board[ni][nj] % 2
		}
	}

	//add & remove it self
	sum -= board[i][j] % 2
	if i-1 >= 0 {
		sum += board[i-1][j] % 2
	}
	return sum
}

func gameOfLife(board [][]int) {
	rowNearby := calc(board, 0, 0)
	for i := range board {
		//new line
		if i > 0 {
			rowNearby = calcRow(board, i, 0, rowNearby)
		}
		nearby := rowNearby

		for j := range board[i] {
			if j > 0 {
				nearby = calcNext(board, i, j, nearby)
			}
			if board[i][j] == 1 && (nearby < 2 || nearby > 3) {
				board[i][j] = 3
			} else if board[i][j] == 0 && nearby == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

func main() {
	board := [][]int{
		// {0, 1, 0},
		// {0, 0, 1},
		// {1, 1, 1},
		// {0, 0, 0},

		// {1, 1},
		// {1, 0},
	}
	gameOfLife(board)

	for i := range board {
		fmt.Println(board[i])
	}
}
