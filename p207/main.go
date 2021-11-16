package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	dependedCourses := make(map[int]int)
	visitedMap := make(map[int]bool)
	queue := make([]int, 0)
	visit := 0
	for i := range prerequisites {
		preMap[prerequisites[i][1]] = append(preMap[prerequisites[i][1]], prerequisites[i][0])
		_, exists := dependedCourses[prerequisites[i][0]]
		if !exists {
			dependedCourses[prerequisites[i][0]] = 1
		} else {
			dependedCourses[prerequisites[i][0]]++
		}
	}
	for i := 0; i < numCourses; i++ {
		_, exists := dependedCourses[i]
		if !exists {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		//dequeue
		q := queue[0]
		queue = queue[1:]
		if _, exists := visitedMap[q]; !exists {
			visit++
			visitedMap[q] = true
			for i := range preMap[q] {
				dependedCourses[preMap[q][i]]--
				if dependedCourses[preMap[q][i]] == 0 {
					queue = append(queue, preMap[q][i])
				}
			}
		}
	}
	return visit == numCourses
}

func main() {
	nums := 4
	pre := [][]int{
		// {1, 0}, {1, 2}, {0, 1},
		{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3},
	}
	fmt.Println(canFinish(nums, pre))
}
