package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	sort.Ints(nums)
	ans := make([]int, len(nums))
	left := (len(nums) - 1) / 2
	right := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ans[i] = nums[left]
			left--
		} else {
			ans[i] = nums[right]
			right--
		}
	}
	for i := range ans {
		nums[i] = ans[i]
	}
}

func main() {
	nums := []int{
		// 1, 5, 1, 1, 6, 4,
		1, 3, 2, 2, 3, 1,
	}
	wiggleSort(nums)
	fmt.Println(nums)
}
