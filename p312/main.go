package main

import "fmt"

func max(x, y int) int {
	if x < y {
		x = y
	}
	return x
}

func maxCoins(nums []int) int {
	size := len(nums)
	state := make([][]int, size+2)
	for i := range state {
		state[i] = make([]int, size+2)
	}
	val := make([]int, size+2)
	val[0], val[size+1] = 1, 1
	for i := 1; i <= size; i++ {
		val[i] = nums[i-1]
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 2; j <= size+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i]*val[k]*val[j] + state[i][k] + state[k][j]
				state[i][j] = max(state[i][j], sum)
			}
		}
	}

	return state[0][size+1]
}

func main() {
	nums := []int{
		3, 1, 5, 8, //167
		// 1, 5, //10
		// 7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3,
	}
	fmt.Println(maxCoins(nums))
}
