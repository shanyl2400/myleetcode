package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
func findDuplicate2(nums []int) int {
	for i := range nums {
		idx := nums[i]
		if idx < 0 {
			idx = -idx
		}
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
			ret := nums[i]
			if nums[i] < 0 {
				ret = -ret
			}
			return ret
		}
	}
	return -1
}

func main() {
	nums := []int{
		// 1, 3, 2, 2, 2, 5,
		3, 1, 3, 4, 2,
	}
	fmt.Println(findDuplicate(nums))
}
