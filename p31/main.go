package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			index = i - 1
			break
		}
	}
	if index < 0 {
		//reverse
		for i := 0; i < len(nums)/2; i++ {
			idx := len(nums) - 1 - i
			temp := nums[i]
			nums[i] = nums[idx]
			nums[idx] = temp
		}
		return
	}
	//find min
	for i := len(nums) - 1; i > index; i-- {
		if nums[i] > nums[index] {
			//swap
			temp := nums[i]
			nums[i] = nums[index]
			nums[index] = temp
			break
		}
	}
	//reverse
	index = index + 1

	sort.Ints(nums[index:])
}

func main() {
	arr := []int{
		3, 2, 1,
	}
	nextPermutation(arr)
	fmt.Println(arr)
}
