package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	size := 1 << len(nums)
	for i := 0; i < size; i++ {
		temp := i
		idx := 0
		ret := make([]int, 0)
		fmt.Println(temp)
		for temp != 0 {
			if temp&0x01 == 1 {
				ret = append(ret, nums[idx])
			}
			idx++
			temp = temp >> 1
		}

		res = append(res, ret)
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1}))
}
