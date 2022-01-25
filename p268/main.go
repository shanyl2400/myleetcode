package main

import "fmt"

func missingNumber(nums []int) int {
	xor := 0
	for i := range nums {
		xor ^= nums[i]
	}
	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}
	return xor
}

func main() {
	nums := []int{
		// 0, 1,
		9, 6, 4, 2, 3, 5, 7, 0, 1,
	}
	fmt.Println(missingNumber(nums))
}
