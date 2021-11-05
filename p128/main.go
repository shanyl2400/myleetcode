package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	max := 0
	for i := range nums {
		numSet[nums[i]] = true
	}

	for i := range nums {
		_, exists := numSet[nums[i]-1]
		if !exists {

			curNum := nums[i]
			streak := 1
			for true {
				_, exists := numSet[curNum+1]
				if !exists {
					break
				}
				streak++
				curNum++
			}
			if max < streak {
				max = streak
			}
		}
	}

	return max
}

func main() {
	nums := []int{
		// 100, 4, 200, 1, 3, 2,
		0, 3, 7, 2, 5, 8, 4, 6, 0, 1,
	}
	fmt.Println(longestConsecutive(nums))
}
