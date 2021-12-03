package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	state := make([][]bool, len(nums))
	for i := range state {
		state[i] = make([]bool, sum+1)
	}
	state[0][nums[0]] = true
	for i := 1; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			state[i][j] = state[i-1][j]

			//从0加到j，不依赖前置条件
			if nums[i] == j {
				state[i][j] = true
				continue
			}

			if nums[i] < j && state[i-1][j-nums[i]] {
				state[i][j] = true
			}
		}
	}
	return state[len(nums)-1][target]
}

func main() {
	nums := []int{
		1, 5, 11, 5,
		// 1, 2, 3, 5,
		// 5, 4, 1, 4, 4,
		// 1,
		// 2, 2, 1, 1,
		// 1, 2, 5,
		// 2, 1, 1,
		// 100,
		// 9, 5,
	}
	fmt.Println(canPartition(nums))
}
