package main

import "fmt"

var dirs = [][]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}
var ans = 0

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	ans = 0
	points := findStart(matrix)

	memo := make([][]int, len(matrix))
	for i := range matrix {
		memo[i] = make([]int, len(matrix[i]))
	}

	for i := range points {
		ans = max(ans, dfs(matrix, points[i][0], points[i][1], memo))
	}
	return ans
}
func findStart(matrix [][]int) [][]int {
	ret := make([][]int, 0)
	for i := range matrix {
		for j := range matrix[i] {
			flag := true
			for _, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				//过滤非法i和j
				if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[i]) {
					continue
				}

				//未访问过，且周围有比matrix[i][j]小的元素，过滤掉
				if matrix[i][j] > matrix[ni][nj] {
					flag = false
				}
			}
			if flag {
				ret = append(ret, []int{i, j})
			}
		}
	}

	return ret
}

func dfs(matrix [][]int, i, j int, memo [][]int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	//访问自身
	memo[i][j]++

	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		//过滤非法i和j
		if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[i]) {
			continue
		}

		//未访问过，且递增
		if matrix[i][j] < matrix[ni][nj] {
			memo[i][j] = max(memo[i][j], dfs(matrix, ni, nj, memo)+1)
		}
	}

	return memo[i][j]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	matrix := [][]int{
		// {9, 9, 4},
		// {6, 6, 8},
		// {2, 1, 1},

		// {3, 4, 5},
		// {3, 2, 6},
		// {2, 2, 1},

		// {1},

		// {1, 2, 3, 4},
		// {15, 15, 12, 5},
		// {11, 12, 13, 6},
		// {10, 9, 8, 7},

		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
		{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
		{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
		{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
		{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
		{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
		{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
		{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
		{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(longestIncreasingPath(matrix))
}
