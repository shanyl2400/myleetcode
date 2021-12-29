package main

import "fmt"

func setZeroes(matrix [][]int) {
	if matrix == nil || matrix[0] == nil {
		return
	}
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}

}
func main() {
	matrix := [][]int{
		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)

	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
