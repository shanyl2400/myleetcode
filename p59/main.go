package main

import "fmt"

var (
	dir = [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
)

func generateMatrix(n int) [][]int {
	if n < 1 {
		return nil
	}
	if n == 1 {
		return [][]int{{1}}
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	size := n * n

	left, right, top, down := 0, n-1, 0, n-1
	idx, dirIdx := 1, 0
	y, x := 0, 0
	for idx <= size {
		matrix[y][x] = idx
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

	return matrix
}

func main() {
	matrix := generateMatrix(5)

	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
