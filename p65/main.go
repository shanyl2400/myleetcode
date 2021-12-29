package main

import "fmt"

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	right := x / 2
	left := 1
	mid := (right + left) / 2
	ans := 1

	for left <= right {
		mid = (right + left) / 2
		if mid*mid-x <= 0 {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	for i := 0; i < 1000; i++ {
		sqrt := mySqrt(i)
		fmt.Println(i, sqrt)
	}
	// fmt.Println(mySqrt(2147395599))
}
