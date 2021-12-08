package main

import "fmt"

func subarraySum(nums []int, k int) int {
	//前缀和
	preSumMap := map[int]int{0: 1}
	preSum := 0
	count := 0
	for i := range nums {
		preSum += nums[i]
		if preSumMap[preSum-k] > 0 {
			count += preSumMap[preSum-k]
		}
		preSumMap[preSum]++
	}

	return count
}

func main() {
	nums := []int{
		// 1, 1, 1,
		1, 2, 3,
	}
	k := 3

	fmt.Println(subarraySum(nums, k))
}
