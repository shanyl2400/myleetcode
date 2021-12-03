package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func largestSumAfterKNegations(nums []int, k int) int {
	freq := make(map[int]int)
	ans := 0
	for i := range nums {
		freq[nums[i]]++
		ans += nums[i]
	}
	for i := -100; i < 0 && k != 0; i++ {
		if freq[i] > 0 {
			ops := min(k, freq[i])
			ans -= i * ops * 2
			freq[-i] += ops
			k -= ops
		}
	}
	if k > 0 && k%2 == 1 && freq[0] == 0 {
		for i := 1; i <= 100; i++ {
			if freq[i] > 0 {
				ans -= i * 2
				break
			}
		}
	}
	return ans
}

func main() {
	nums := []int{
		// 4, 2, 3,
		3, -1, 0, 2,
		// 2, -3, -1, 5, -4,
	}
	k := 3
	fmt.Println(largestSumAfterKNegations(nums, k))
}
