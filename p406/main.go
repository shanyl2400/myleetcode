package main

import "fmt"

func updateState(start int, idx int, state []int, people [][]int) {
	for i := start; i < len(people); i++ {
		if i == idx {
			continue
		}
		if people[i][0] <= people[idx][0] {
			state[i]++
		}
	}
}

func reconstructQueue(people [][]int) [][]int {
	state := make([]int, len(people))
	for idx := range people {
		//排序第idx个people

		picked := -1
		//选择
		for j := idx; j < len(people); j++ {
			//符合条件
			if people[j][1] == state[j] {
				//选取最小的值
				if picked < 0 || people[picked][0] > people[j][0] {
					picked = j
				}
			}
		}
		//没有合适的
		if picked < 0 {
			fmt.Println("no valid")
			return people
		}
		//更新状态
		updateState(idx, picked, state, people)
		state[picked] = state[idx]

		//排序并更新状态
		people[idx], people[picked] = people[picked], people[idx]
		// state[idx], state[picked] = state[picked], state[idx]
	}

	return people
}

func main() {
	people := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	// people := [][]int{
	// 	{6, 0},
	// 	{5, 0},
	// 	{4, 0},
	// 	{3, 2},
	// 	{2, 2},
	// 	{1, 4},
	// }

	fmt.Println(reconstructQueue(people))
}
