package main

import "fmt"

func helper(n int) int {
	ret := 0
	for ; n > 0; n /= 10 {
		tmp := n % 10
		ret += tmp * tmp
	}
	return ret
}

func isHappy(n int) bool {
	his := make(map[int]bool)
	for n != 1 {
		his[n] = true
		n = helper(n)
		if his[n] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isHappy(2))
}
