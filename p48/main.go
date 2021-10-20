package main

import "fmt"

func rotate(matrix [][]int) {
	rows := len(matrix)
	if rows < 1 {
		return
	}
	cols := len(matrix[0])

	output := make([][]int, rows)
	for i := range output {
		output[i] = make([]int, cols)
	}

	for i := range matrix {
		for j := range matrix[i] {
			output[i][rows-j-1] = matrix[j][i]
		}
	}
	for i := range output {
		for j := range output[i] {
			matrix[i][j] = output[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(matrix)

	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
