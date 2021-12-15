package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	left, right := 0, 1
	for ; right < len(nums); right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	fmt.Println(nums, left)
	return left + 1
}

func main() {
	nums := []int{
		// 0, 0, 1, 1, 2, 2, 3, 3, 4,
		// 1, 1, 2,
	}
	fmt.Println(removeDuplicates(nums))
}
