package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	state := make([][]int, m)
	for i := range state {
		state[i] = make([]int, n)
	}
	state[0][0] = grid[m-1][n-1]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				state[i][j] = grid[m-i-1][n-j-1] + state[i][j-1]
			} else if j == 0 {
				state[i][j] = grid[m-i-1][n-j-1] + state[i-1][j]
			} else {
				state[i][j] = grid[m-i-1][n-j-1] + min(state[i-1][j], state[i][j-1])
			}
		}
	}

	return state[m-1][n-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	// grid := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	fmt.Println(minPathSum(grid))
}
