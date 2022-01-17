package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) < 1 {
		return ""
	}
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10

		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	ans := []byte{}
	sum := 0
	for _, x := range nums {
		sum += x
		ans = append(ans, strconv.Itoa(x)...)
	}
	if sum == 0 {
		return "0"
	}
	return string(ans)
}

func main() {
	nums := []int{
		// 3, 30, 34, 5, 9,
		// 10, 2,
		// 1,
		// 10,
		// 34323, 3432,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
		// 12341, 123411234, //12341123412341
		//12341123411234
		//12341123412341
	}
	fmt.Println(largestNumber(nums))
}
