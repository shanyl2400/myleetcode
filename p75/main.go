package main

import "fmt"

func sortColors(nums []int) {
	p0 := 0 // pointer 0
	p1 := 0 // pointer 1

	for idx := range nums {
		if nums[idx] == 0 {
			swap(nums, p0, idx)
			if p0 != p1 {
				swap(nums, p1, idx)
			}
			p0++
			p1++
		} else if nums[idx] == 1 {
			swap(nums, p1, idx)
			p1++
		}
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {
	nums := []int{
		2, 0, 2, 1, 1, 0,
		// 2, 0, 1,
		// 0,
		// 1,
		// 1, 1, 2, 0, 0, 1, 2, 2, 0, 0, 1, 1,
		// 1, 1, 2, 0, 0, 1, 2,
	}
	sortColors(nums)
	fmt.Println(nums)
}
