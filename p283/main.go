package main

import "fmt"

func moveZeroes(nums []int) {
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{
		0, 1, 0, 3, 12, 0,
	}
	moveZeroes(nums)
	fmt.Println(nums)
}
