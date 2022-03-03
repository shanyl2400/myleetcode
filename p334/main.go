package main

import (
	"fmt"
	"math"
)

func increasingTriplet2(nums []int) bool {
	leftMin := make([]int, len(nums))
	rightMax := make([]int, len(nums))

	leftMin[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if leftMin[i-1] > nums[i] {
			leftMin[i] = nums[i]
		} else {
			leftMin[i] = leftMin[i-1]
		}
	}

	rightMax[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if rightMax[i+1] < nums[i] {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32

	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else if nums[i] < first {
			first = nums[i]
		}
	}
	return false
}
func main() {
	nums := []int{
		2, 1, 5, 0, 4, 6,
	}
	fmt.Println(increasingTriplet(nums))
}
