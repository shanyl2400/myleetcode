package main

import "fmt"

func hammingDistance(x int, y int) int {
	target := x ^ y
	ret := 0
	for target != 0 {
		ret++
		target = target & (target - 1)
	}

	return ret
}

func main() {
	fmt.Println(hammingDistance(1, 3))
}
