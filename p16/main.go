package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := 0
	minDiff := 0
	first := true

	start := 0
	end := 0
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		start = i + 1
		end = len(nums) - 1
		newTarget := target - nums[i]
		for start < end {
			fmt.Println("i:", i, "start:", start, "end: ", end)
			diff := newTarget - nums[start] - nums[end]
			absDiff := diff
			if absDiff < 0 {
				absDiff = -diff
			}
			fmt.Println("newtarget:", newTarget, "minDiff:", minDiff, "absDiff:", absDiff)
			fmt.Println("sum:", nums[start]+nums[end]+nums[i])
			if first || absDiff < minDiff {
				minDiff = absDiff
				sum = nums[start] + nums[end] + nums[i]
				first = false
			}
			fmt.Println("diff:", diff)

			if diff < 0 {
				end--
			} else if diff > 0 {
				start++
			} else {
				return target
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(threeSumClosest(
		[]int{1, 1, 1, 0},
		// []int{-1, 2, 1, -4},
		// []int{0, 2, 1, -3},
		1,
	))
}
