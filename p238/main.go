package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	R := 1
	for i := len(nums) - 2; i >= 0; i-- {
		// R = R * nums[i]
		R = R * nums[i+1]
		fmt.Println(R)
		answer[i] = answer[i] * R
	}
	// answer[0] = R

	return answer
}

func main() {
	nums := []int{
		1, 2, 3, 4,
	}
	fmt.Println(productExceptSelf(nums))
}
