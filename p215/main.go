package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	stack1 := make([]int, len(nums))
	stack2 := make([]int, len(nums))
	top1 := -1
	top2 := -1

	for i := range nums {
		if top1 < 0 {
			top1++
			stack1[top1] = nums[i]
			continue
		}

		if nums[i] > stack1[top1] || top1 < k-1 {
			for top1 >= 0 && nums[i] > stack1[top1] {
				//reset stack
				top2++
				stack2[top2] = stack1[top1]
				top1--
			}
			top1++
			stack1[top1] = nums[i]
			for top2 >= 0 && top1 < k-1 {
				top1++
				stack1[top1] = stack2[top2]
				top2--
			}
			top2 = -1
		}

	}
	return stack1[top1]
}

func main() {
	nums := []int{
		// 3, 2, 1, 5, 6, 4,
		// 3, 2, 3, 1, 2, 4, 5, 5, 6,
		2, 1, 0,
	}
	fmt.Println(findKthLargest(nums, 2))
}
