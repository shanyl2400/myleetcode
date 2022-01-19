package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	degrees := make([]int, numCourses)
	connMap := make(map[int][]int)
	for i := range prerequisites {
		degrees[prerequisites[i][0]]++
		connMap[prerequisites[i][1]] = append(connMap[prerequisites[i][1]], prerequisites[i][0])
	}
	queue := make([]int, 0)
	for i := range degrees {
		if degrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	ans := make([]int, 0)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for _, val := range connMap[item] {
			degrees[val]--
			if degrees[val] == 0 {
				queue = append(queue, val)
			}
		}

		ans = append(ans, item)
	}
	if len(ans) != numCourses {
		return nil
	}
	return ans
}

func main() {
	// numCourses := 4
	// prerequisites := [][]int{
	// 	{1, 0},
	// 	{2, 0},
	// 	{3, 1},
	// 	{3, 2},
	// }

	// numCourses := 10
	// prerequisites := [][]int{
	// 	{1, 0},
	// 	{1, 5},
	// 	{2, 1},
	// 	{3, 2},
	// 	{3, 7},
	// 	{8, 3},
	// 	{8, 7},
	// }
	numCourses := 3
	prerequisites := [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}

	fmt.Println(findOrder(numCourses, prerequisites))
}
