package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if target < 0 {
		target = -target
	}
	if target > sum {
		return 0
	}

	state := make([][]int, 2)
	for i := range state {
		state[i] = make([]int, sum+1)
	}

	state[0][nums[0]] = 1
	if nums[0] == 0 {
		state[0][nums[0]] = 2
	}
	for i := 1; i < len(nums); i++ {
		prv := (i + 1) % 2
		cur := i % 2
		for j := 0; j <= sum; j++ {
			if state[prv][j] > 0 {
				//(j & -j) +_ nums[i]
				if nums[i]+j >= 0 {
					state[cur][nums[i]+j] += state[prv][j]
				}
				if nums[i]-j >= 0 && j != 0 {
					state[cur][nums[i]-j] += state[prv][j]
				}
				if -nums[i]+j >= 0 {
					state[cur][-nums[i]+j] += state[prv][j]
				}
				if -nums[i]-j >= 0 && j != 0 {
					state[cur][-nums[i]-j] += state[prv][j]
				}
			}
		}
		for j := 0; j <= sum; j++ {
			state[prv][j] = 0
		}
	}

	return state[(len(nums)-1)%2][target]
}

func main() {
	nums := []int{
		1, 1, 1, 1, 1,
		// 1,
		// 0, 0, 0, 0, 0, 0, 0, 0, 1,
		// 100,
	}
	target := 3

	fmt.Println(findTargetSumWays(nums, target))
}
