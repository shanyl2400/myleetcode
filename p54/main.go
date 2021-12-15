package main

import "fmt"

var (
	dir = [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return nil
	}
	rows := len(matrix)
	cols := len(matrix[0])

	left, right, top, down := 0, cols-1, 0, rows-1
	idx, dirIdx := 0, 0
	y, x := 0, 0
	ret := make([]int, cols*rows)
	for idx < cols*rows {
		ret[idx] = matrix[y][x]
		if x == right && dirIdx == 0 {
			top++
			dirIdx = (dirIdx + 1) % len(dir)
		} else if x == left && dirIdx == 2 {
			down--
			dirIdx = (dirIdx + 1) % len(dir)
		} else if y == down && dirIdx == 1 {
			right--
			dirIdx = (dirIdx + 1) % len(dir)
		} else if y == top && dirIdx == 3 {
			left++
			dirIdx = (dirIdx + 1) % len(dir)
		}

		y += dir[dirIdx][0]
		x += dir[dirIdx][1]
		idx++
	}

	return ret
}

func main() {
	matrix := [][]int{
		// {1, 2, 3},
		// {4, 5, 6},
		// {7, 8, 9},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
