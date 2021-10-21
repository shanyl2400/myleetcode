package main

import "fmt"

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
	ret := combinationSum(candidates, 9)
	// for i := range ret {
	// 	fmt.Println(ret[i])
	// }
	fmt.Println(ret)
}

// [2,2,2,3]

// [2 2 2 1]
