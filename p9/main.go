package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	palind := 0
	for temp != 0 {
		d := temp % 10
		palind = palind * 10
		palind = palind + d

		temp = temp / 10
	}
	return palind == x
}
func main() {
	fmt.Println(isPalindrome(-101))
}
