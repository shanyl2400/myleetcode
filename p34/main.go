package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{
			-1, -1,
		}
	}
	first := findFirst(nums, target)
	if first == -1 {
		return []int{
			-1, -1,
		}
	}
	second := findLast(nums, target)

	return []int{first, second}
}
func findFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			right = mid
		} else {
			right = mid - 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func findLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2 //向上取整
		if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	nums := []int{
		// 5, 7, 7, 8, 8, 10,
	}
	fmt.Println(searchRange(nums, 6))
}
