package main

import "fmt"

func findPeakElement(nums []int) int {
	idx := 0
	for i := range nums {
		if nums[idx] < nums[i] {
			idx = i
		}
	}
	return idx
}

func main() {
	nums := []int{
		1, 2, 1, 3, 5, 6, 4,
	}
	fmt.Println(findPeakElement(nums))
}
