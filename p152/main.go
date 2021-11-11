package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	state := make([]int, len(nums))
	minState := make([]int, len(nums))
	state[0] = nums[0]
	minState[0] = nums[0]
	ret := state[0]
	for i := 1; i < len(nums); i++ {
		data1 := state[i-1] * nums[i]
		data2 := minState[i-1] * nums[i]
		state[i] = max(data1, data2, nums[i])
		minState[i] = min(data1, data2, nums[i])

		if ret < state[i] {
			ret = state[i]
		}
	}
	fmt.Println(state)
	fmt.Println(minState)
	return ret
}

func max(x, y, z int) int {
	if x < y {
		x = y
	}
	if x < z {
		x = z
	}
	return x
}

func min(x, y, z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

func main() {
	nums := []int{
		2, 3, -2, 4,
		// -2, 0, -1,
		// -3, -1, -1,
		// -1, -2, -9, -6,
	}
	fmt.Println(maxProduct(nums))
}
