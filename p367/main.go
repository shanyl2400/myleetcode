package main

import "fmt"

func isPerfectSquare(num int) bool {
	start := 0
	end := num

	for start <= end {
		mid := (start + end) >> 1
		x := mid * mid
		if x == num {
			return true
		} else if x > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
		fmt.Println(start, end)
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(14))
}
