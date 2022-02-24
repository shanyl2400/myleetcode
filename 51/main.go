package main

import "fmt"

func reversePairs(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	temp := make([]int, len(nums))
	return reverse(nums, 0, len(nums)-1, temp)
}

func reverse(nums []int, left, right int, temp []int) int {
	if left == right {
		return 0
	}
	mid := (left + right) / 2
	leftPart := reverse(nums, left, mid, temp)
	rightPart := reverse(nums, mid+1, right, temp)

	if nums[mid] <= nums[mid+1] {
		return leftPart + rightPart
	}

	return merge(nums, left, mid, right, temp) + leftPart + rightPart
}
func merge(nums []int, left, mid, right int, temp []int) int {
	i, j := left, mid+1
	count := 0
	for k := 0; k < right-left+1; k++ {
		if i == mid+1 {
			temp[k] = nums[j]
			j++
		} else if j == right+1 {
			temp[k] = nums[i]
			i++
		} else if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++

			count += (mid - i + 1)
		}
	}
	for i := left; i <= right; i++ {
		nums[i] = temp[i-left]
	}
	return count
}
func main() {
	nums := []int{
		7, 5, 6, 4,
		// 8, 64, 12, 9, 16, 55, 22, 100, 26, 91,
	}
	fmt.Println(reversePairs(nums))
}
