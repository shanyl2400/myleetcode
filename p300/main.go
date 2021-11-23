package main

import "fmt"

func lengthOfLIS(nums []int) int {
	ret := 0
	state := make([]int, len(nums))
	for i := range nums {
		state[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			if state[j]+1 > state[i] {
				state[i] = state[j] + 1
			}
		}
		if ret < state[i] {
			ret = state[i]
		}
	}
	return ret
}
func main() {
	nums := []int{
		10, 9, 2, 5, 3, 7, 101, 18,
		// 0, 1, 0, 3, 2, 3,
		// 7, 7, 7, 7, 7, 7, 7,
	}
	fmt.Println(lengthOfLIS(nums))
}
