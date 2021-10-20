package main

import "fmt"

func divide(dividend int, divisor int) int {
	neg := false
	if dividend < 0 {
		dividend = -dividend
		neg = !neg
	}
	if divisor < 0 {
		divisor = -divisor
		neg = !neg
	}

	bitSize := 0
	temp := dividend
	for temp != 0 {
		bitSize++
		temp = temp >> 1
	}
	rest := 0
	out := 0
	for i := 0; i < bitSize; i++ {
		bit := dividend >> (bitSize - i - 1) & 1
		rest = rest<<1 + bit
		out = out << 1
		if rest >= divisor {
			rest = rest - divisor
			out = out + 1
		}
		// fmt.Println(rest)
	}
	if neg {
		out = -out
	}
	if out < -2147483648 {
		out = -2147483648
	} else if out > 2147483647 {
		out = 2147483647
	}
	return out
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
