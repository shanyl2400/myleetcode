package main

import "fmt"

func majorityElement(nums []int) []int {
	vote1, vote2 := 0, 0
	element1, element2 := 0, 0

	for i := range nums {
		if vote1 > 0 && element1 == nums[i] {
			vote1++
		} else if vote2 > 0 && element2 == nums[i] {
			vote2++
		} else if vote1 == 0 {
			element1 = nums[i]
			vote1++
		} else if vote2 == 0 {
			element2 = nums[i]
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	cnt1, cnt2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == element1 {
			cnt1++
		} else if vote2 > 0 && nums[i] == element2 {
			cnt2++
		}
	}
	ret := make([]int, 0)
	if vote1 > 0 && cnt1 > len(nums)/3 {
		ret = append(ret, element1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ret = append(ret, element2)
	}
	return ret
}

func main() {
	nums := []int{
		3, 2, 3,
	}
	fmt.Println(majorityElement(nums))
}
