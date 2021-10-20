package main

import "fmt"

func search(nums []int, target int) int {

	return findNum(nums, target, 0, len(nums)-1)
}
func findNum(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	} else if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}

	//左边有序
	if nums[start] < nums[mid] {
		if target < nums[mid] && target > nums[start] {
			return findNum(nums, target, start, mid-1)
		} else {
			return findNum(nums, target, mid+1, end)
		}
	} else {
		//右边有序
		if target > nums[mid] && target < nums[end] {
			return findNum(nums, target, mid+1, end)
		} else {
			return findNum(nums, target, start, mid-1)
		}
	}
}

func main() {
	nums := []int{
		4, 5, 6, 7, 0, 1, 2,
		// 4, 5, 6, 7, 0, 1, 2,
		// 1,
	}
	fmt.Println(search(nums, 0))
}
