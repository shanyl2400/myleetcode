package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfs(candidates, target, &res, []int{}, 0)
	return res
}

func dfs(candidates []int, target int, res *[][]int, currentList []int, idx int) {
	if idx >= len(candidates) {
		return
	}
	if target == 0 {
		list := make([]int, len(currentList))
		copy(list, currentList)
		*res = append(*res, list)
		return
	}
	dfs(candidates, target, res, currentList, idx+1)

	if candidates[idx] <= target {
		currentList = append(currentList, candidates[idx])
		dfs(candidates, target-candidates[idx], res, currentList, idx)
		currentList = currentList[:len(currentList)-1]
	}
}
func combinationSum(candidates []int, target int) [][]int {
	state := make([][][]int, target+1)
	for i := range candidates {
		if candidates[i] <= target {
			state[target-candidates[i]] = append(state[target-candidates[i]], []int{candidates[i]})
		}
		for j := 1; j <= target; j++ {
			sub := j - candidates[i]
			times := 1
			for sub >= 0 {
				for x := range state[j] {
					temp := make([]int, len(state[j][x]))
					copy(temp, state[j][x])
					for p := 0; p < times; p++ {
						temp = append(temp, candidates[i])
					}
					state[sub] = append(state[sub], temp)
				}
				times++
				sub = j - candidates[i]*times
			}
		}
	}
	return state[0]
}
func main() {
	candidates := []int{
		// 2, 3, 6, 7,
		// 2, 3, 5,
		2, 7, 6, 3, 5, 1,
		// 1,
	}
	ret := combinationSum2(candidates, 9)
	// for i := range ret {
	// 	fmt.Println(ret[i])
	// }
	fmt.Println(ret)
}
