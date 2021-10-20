package main

import "fmt"

func firstMissingPositive(nums []int) int {
	flag := false
	for i := range nums {

		if nums[i] == 1 {
			flag = true
		} else if nums[i] < 1 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	if !flag {
		return 1
	}
	for _, v := range nums {
		index := v
		if v < 0 {
			index = -index
		}

		x := nums[index-1]
		if x > 0 {
			nums[index-1] = -x
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{
		7, 8, 9, 11, 12,
	}))
}
