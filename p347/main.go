package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	max := 0
	for i := range nums {
		freq[nums[i]]++
		f := freq[nums[i]]
		if f > max {
			max = f
		}
	}
	counter := make([][]int, max+1)
	for k, v := range freq {
		counter[v] = append(counter[v], k)
	}
	ret := make([]int, 0, k)
	for i := max; i >= 0; i-- {
		ret = append(ret, counter[i]...)
		if len(ret) == k {
			break
		}
	}

	return ret
}

func main() {
	nums := []int{
		1, 1, 1, 2, 2, 3,
		// 1, 2,
	}
	fmt.Println(topKFrequent(nums, 2))
}
