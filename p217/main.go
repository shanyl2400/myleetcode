package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}

func main() {
	nums := []int{
		// 1, 2, 3, 1,
		// 1, 2, 3, 4,
		1, 1, 1, 3, 3, 4, 3, 2, 4, 2,
	}
	fmt.Println(containsDuplicate(nums))
}
