package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	idx := 0
	for i := range nums {
		idx = nums[i]
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	res := make([]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	nums := []int{
		// 4, 3, 2, 7, 8, 2, 3, 1,
		1, 1,
	}
	fmt.Println(findDisappearedNumbers(nums))
}
