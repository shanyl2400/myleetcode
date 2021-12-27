package main

import "fmt"

const (
	rowOffset  = 0
	colOffset  = 9
	cellOffset = 18
)

func isValidSudoku(board [][]byte) bool {
	record := make([]uint16, 27)
	for y := range board {
		for x := range board {
			val := board[y][x]
			if val == '.' {
				continue
			}
			val = val - '0'

			row, col := y, x
			cell := x/3 + 3*(y/3)
			//check val
			//check row
			d := uint16(1 << (val - 1))
			if record[rowOffset+row]&d != 0 {
				return false
			}
			record[rowOffset+row] = record[rowOffset+row] | d

			//check col
			if record[colOffset+col]&d != 0 {
				return false
			}
			record[colOffset+col] = record[colOffset+col] | d

			//check cell
			if record[cellOffset+cell]&d != 0 {
				return false
			}
			record[cellOffset+cell] = record[cellOffset+cell] | d
			fmt.Println("val:", int(val), "row:", row, "col:", col, "cell:", cell)
			for i := range record {
				fmt.Printf("%09b ", record[i])
			}
			fmt.Println()
		}
	}
	return true
}

func main() {
	// board := [][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

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
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '9', '.', '.', '.', '.', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '9', '9', '3', '5', '7', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
		{'.', '.', '.', '8', '.', '.', '.', '.', '.'},
		{'.', '1', '.', '.', '.', '.', '4', '.', '9'},
		{'.', '.', '.', '5', '.', '4', '.', '.', '.'},
	}

	fmt.Println(isValidSudoku(board))
}
