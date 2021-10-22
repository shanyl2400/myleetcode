package main

import "fmt"

func plusOne(digits []int) []int {
	closed := false
	for i := len(digits) - 1; i >= 0; i-- {
		x := digits[i] + 1
		if x > 9 {
			digits[i] = 0
		} else {
			digits[i] = x
			closed = true
			break
		}
	}
	if !closed {
		return append([]int{1}, digits...)
	}
	fmt.Println(closed)
	return digits
}

func main() {
	fmt.Println(plusOne([]int{
		// 1, 2, 3,
		// 4, 3, 2, 2,
		0,
		// 9, 9, 9, 9,
	}))
}
