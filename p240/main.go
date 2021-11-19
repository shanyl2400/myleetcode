package main

import "fmt"

type pos struct {
	i int
	j int
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if target > matrix[x][y] {
			x++
		} else {
			y--
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	for i := range matrix {
		end := len(matrix[i]) - 1
		start := 0
		nums := matrix[i]
		if nums[start] > target || nums[end] < target {
			continue
		}
		//二分查找
		for start <= end {
			mid := (start + end) / 2
			if nums[mid] == target {
				return true
			}
			if nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return false
}

func searchMatrix3(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	visited := make([][]bool, len(matrix))
	for i := range matrix {
		visited[i] = make([]bool, len(matrix[i]))
	}
	size := len(matrix)*len(matrix[0]) + 1

	if matrix[0][0] > target {
		return false
	}
	queue := make([]pos, size)
	front := 0
	back := 0

	queue[back] = pos{i: 0, j: 0}
	visited[0][0] = true

	for front <= back {
		p := queue[front]
		front++
		// queue = queue[1:]

		if matrix[p.i][p.j] == target {
			return true
		}

		if p.i+1 < len(matrix) && matrix[p.i+1][p.j] <= target && !visited[p.i+1][p.j] {
			// queue = append(queue, pos{i: p.i + 1, j: p.j})
			back++
			visited[p.i+1][p.j] = true
			queue[back] = pos{i: p.i + 1, j: p.j}
		}
		if p.j+1 < len(matrix[p.i]) && matrix[p.i][p.j+1] <= target && !visited[p.i][p.j+1] {
			// queue = append(queue, pos{i: p.i, j: p.j + 1})
			back++
			visited[p.i][p.j+1] = true
			queue[back] = pos{i: p.i, j: p.j + 1}
		}
	}
	return false
}
func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5

	fmt.Println(searchMatrix(matrix, target))
}
