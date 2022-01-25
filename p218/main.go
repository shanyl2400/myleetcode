package main

import "fmt"

func getSkyline(buildings [][]int) [][]int {
	max := 0
	for i := range buildings {
		if max < buildings[i][1] {
			max = buildings[i][1]
		}
	}
	arr := make([]int, max+1)

	for i := range buildings {
		for j := buildings[i][0]; j < buildings[i][1]; j++ {
			if arr[j] < buildings[i][2] {
				arr[j] = buildings[i][2]
			}
		}
	}
	prev := 0
	ans := make([][]int, 0)
	for i := range arr {
		if arr[i] != prev {
			ans = append(ans, []int{
				i, arr[i],
			})
			prev = arr[i]
		}
	}
	return ans
}

func main() {
	building := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	// building := [][]int{
	// 	{0, 2, 3},
	// 	{2, 5, 3},
	// }
	fmt.Println(getSkyline(building))
}
