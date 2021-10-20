package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	R := 0
	//排序
	sort.Ints(nums)
	ret := make([][]int, 0)
	// fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		R = len(nums) - 1
		sum := 1
		for L := i + 1; L < len(nums); L++ {
			if L > i+1 && nums[L] == nums[L-1] {
				continue
			}
			sum = nums[i] + nums[L] + nums[R]
			for L < R && sum > 0 {
				R--
				sum = nums[i] + nums[L] + nums[R]
			}
			if L == R {
				break
			}

			if sum == 0 {
				ret = append(ret, []int{
					nums[L], nums[i], nums[R],
				})
				R = len(nums) - 1
				// break
			}
			// fmt.Println("sum:", sum, "L:", L, "R:", R, "i:", i)

		}
	}
	//去重

	return ret
}
func main() {
	ret := threeSum([]int{
		// -1, 0, 1, 2, -1, -4,
		0, 0, 0, 0,
	})
	for i := range ret {
		fmt.Println(ret[i])
	}
}
