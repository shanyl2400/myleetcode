package main

import "fmt"

func countSmaller(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	indexes := make([]int, len(nums))
	for i := range indexes {
		indexes[i] = i
	}
	temp := make([]int, len(nums))
	res := make([]int, len(nums))
	reverse(nums, 0, len(nums)-1, indexes, res, temp)
	return res
}

func reverse(nums []int, left, right int, indexes, res, temp []int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	reverse(nums, left, mid, indexes, res, temp)
	reverse(nums, mid+1, right, indexes, res, temp)

	if nums[indexes[mid]] <= nums[indexes[mid+1]] {
		return
	}

	merge(nums, left, mid, right, indexes, res, temp)
}
func merge(nums []int, left, mid, right int, indexes, res, temp []int) {
	for i := left; i <= right; i++ {
		temp[i] = indexes[i]
	}
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		if i > mid {
			indexes[k] = temp[j]
			j++
		} else if j > right {
			indexes[k] = temp[i]
			i++
			res[indexes[k]] += right - mid
		} else if nums[temp[i]] <= nums[temp[j]] {
			indexes[k] = temp[i]
			i++
			res[indexes[k]] += (j - mid - 1)
		} else {
			indexes[k] = temp[j]
			j++
		}
	}
}

func main() {
	nums := []int{
		5, 2, 6, 1,
		// -1, -1,
	}
	fmt.Println(countSmaller(nums))
}
