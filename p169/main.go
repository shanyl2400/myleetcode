package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	cur := 0
	for i := range nums {
		if count < 1 {
			cur = nums[i]
			count = 1
			continue
		}
		if nums[i] == cur {
			count++
		} else {
			count--
		}
	}
	return cur
}

func main() {
	nums := []int{
		3, 3, 4,
	}
	fmt.Println(majorityElement(nums))
}
