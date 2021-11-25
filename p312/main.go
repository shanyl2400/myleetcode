package main

import "fmt"

func getCoins(nums []int, pick int) int {
	x := 1
	y := nums[pick]
	z := 1
	if pick > 0 {
		x = nums[pick-1]
	}
	if pick < len(nums)-1 {
		z = nums[pick+1]
	}
	return x * y * z
}

func copyNums(nums []int, exclude int) []int {
	newNum := make([]int, 0, len(nums)-1)
	for i := range nums {
		if i != exclude {
			newNum = append(newNum, nums[i])
		}
	}
	return newNum
}

func maxCoins(nums []int) int {
	max := 0
	for i := range nums {
		res := getCoins(nums, i)
		if len(nums) > 1 {
			nn := copyNums(nums, i)

			res = res + maxCoins(nn)
		}

		if max < res {
			max = res
		}
	}
	return max
}

func main() {
	nums := []int{
		3, 1, 5, 8, //167
		// 1, 5, //10
		// 7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3,
	}
	fmt.Println(maxCoins(nums))
}
