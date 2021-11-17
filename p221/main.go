package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				matrix[i][j] = 0
			} else {
				if i == 0 || j == 0 {
					matrix[i][j] = 1
				} else {
					matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1]) + 1
				}
				if max < int(matrix[i][j]) {
					max = int(matrix[i][j])
				}
			}
		}
	}
	return max * max
}
func min(x, y, z byte) byte {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

func main() {
	matrix := [][]byte{
		// {'1', '0', '1', '0', '0'},
		// {'1', '0', '1', '1', '1'},
		// {'1', '1', '1', '1', '1'},
		// {'1', '0', '0', '1', '0'},
		{'0', '1'},
		{'1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
