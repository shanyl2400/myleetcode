package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	//进一法获取中间值
	rest := size % 2
	mid := size / 2

	index1 := 0
	index2 := 0
	val1 := 0
	val2 := 0
	index := index1 + index2
	for index != mid+1 {
		if len(nums1) <= index1 {
			val1 = nums2[index2]
			index2++
		} else if len(nums2) <= index2 {
			val1 = nums1[index1]
			index1++
		} else if nums1[index1] < nums2[index2] {
			val1 = nums1[index1]
			index1++
		} else {
			val1 = nums2[index2]
			index2++
		}
		index = index1 + index2
		if rest == 0 && index == mid {
			val2 = val1
		}
	}
	if rest == 1 {
		return float64(val1)
	}
	return float64(val1+val2) / 2
}