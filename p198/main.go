package main

import "fmt"

func rob(nums []int) int {
	state := make([]int, len(nums))
	for i := range nums {
		tmp := nums[i]
		for j := 0; j <= i-2; j++ {
			d := state[j] + nums[i]
			if tmp < d {
				tmp = d
			}
		}
		state[i] = tmp
	}
	max := state[0]
	for i := range state {
		if max < state[i] {
			max = state[i]
		}
	}
	return max
}

func main() {
	nums := []int{
		// 1, 2, 3, 1,
		// 2, 1, 1, 2,
		// 2, 7, 9, 3, 1,
		114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240,
	}
	fmt.Println(rob(nums))
}
