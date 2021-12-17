package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2 := m-1, n-1
	idx := m + n - 1

	for idx1 >= 0 || idx2 >= 0 {
		if idx1 < 0 || (idx2 >= 0 && nums1[idx1] < nums2[idx2]) {
			nums1[idx] = nums2[idx2]
			idx2--
		} else {
			nums1[idx] = nums1[idx1]
			idx1--
		}
		idx--
	}
}

func main() {
	nums1 := []int{
		// 1, 2, 3, 0, 0, 0,
		// 1,
		0,
	}
	nums2 := []int{
		// 2, 5, 6,
		1,
	}
	// m, n := 3, 3
	m, n := 0, 1
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
