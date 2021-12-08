package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
