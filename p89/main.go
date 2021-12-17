package main

import (
	"fmt"
	"math"
)

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	size := int(math.Pow(2, float64(n)))
	ret := make([]int, 0, size)
	for num := 0; num < size; num++ {
		code := 0
		base := 1
		for bit := 0; bit < n; bit++ {
			tag := 1 << bit
			numMode := num % (tag * 4)
			if numMode >= base && numMode < base+base*2 {
				code |= tag
			} else {
				code &= ^tag
			}
			base = base * 2
		}
		ret = append(ret, code)
	}

	return ret
}

func main() {
	fmt.Println(grayCode(3))
}
