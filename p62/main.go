package main

import "fmt"

func uniquePaths(m int, n int) int {
	state := make([][]int, m)
	for i := range state {
		state[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		state[i][0] = 1
	}
	for j := 1; j < n; j++ {
		state[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			state[i][j] = state[i-1][j] + state[i][j-1]
		}
	}

	return state[m-1][n-1]
}
func uniquePaths2(m int, n int) int {
	return doUniquePaths(m-1, n-1)
}
func doUniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	res := 0
	if m > 0 {
		res = res + doUniquePaths(m-1, n)
	}
	if n > 0 {
		res = res + doUniquePaths(m, n-1)
	}
	return res
}

func main() {
	fmt.Println(uniquePaths(3, 3))
}
