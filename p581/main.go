package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	size := len(nums)
	l, r := 0, size-1
	min, max := nums[r], nums[l]
	for i := range nums {
		if nums[i] >= max {
			max = nums[i]
		} else {
			l = i
		}

		if nums[size-i-1] <= min {
			min = nums[size-i-1]
		} else {
			r = size - i - 1
		}
	}
	if l == 0 && r == size-1 {
		return 0
	}
	return l - r + 1
}

func main() {
	nums := []int{
		// 2, 6, 4, 8, 10, 9, 15,
		1, 2, 3, 4,
	}
	fmt.Println(findUnsortedSubarray(nums))
}
